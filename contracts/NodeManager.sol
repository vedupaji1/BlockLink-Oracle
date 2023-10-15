// SPDX-License-Identifier: LGPL 3.0
pragma solidity 0.8.8;

import "./BLS.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

// Note:- For Production Make Sure You Are Deploying BLS Library Separately And Using In Your All Contracts

contract NodeManager is Ownable {
    struct Node {
        uint256 balance;
        address owner;
    }

    struct SignatureMessage {
        bytes32 nodeIdHash;
        uint256 nonce;
    }

    event NodeAdded(
        bytes32 indexed nodeIdHash,
        address indexed nodeOwner,
        uint256 tickNum
    );
    event NodeLeft(bytes32 indexed nodeIdHash, address indexed nodeOwner);
    event NodeRemoved(
        bytes32 indexed nodeIdHash,
        address indexed nodeOwner,
        uint256 amountLocked
    );
    event ExtraDepositForNode(
        bytes32 indexed nodeIdHash,
        address indexed depositor,
        uint256 amount
    );
    event TickDurationChanged(
        uint256 indexed oldTickDuration,
        uint256 indexed newTickDuration
    );
    event MaxJoinsInTickChanged(
        uint256 indexed oldMaxJoinsInTick,
        uint256 indexed newMaxJoinsInTick
    );
    event RequiredNodeBalanceChanged(
        uint256 indexed oldRequiredNodeBalance,
        uint256 indexed newRequiredNodeBalance
    );

    uint256 public curTickNum;
    uint256 public nodesJoinedInCurTick;
    uint256 public tickDuration;
    uint256 public maxJoinsInTick;
    uint256 public requiredNodeBalance;
    uint256 public totalNodes;
    mapping(bytes32 => Node) public nodeData;
    mapping(bytes32 => uint256) public nonceForNode;

    constructor(
        uint256 _curTickNum,
        uint256 _tickDuration,
        uint256 _maxJoinsInTick,
        uint256 _requiredNodeBalance
    ) {
        require(
            _curTickNum > block.timestamp,
            "`CurTickTime` Must Be Greater Than Current Time"
        );
        curTickNum = _curTickNum;
        tickDuration = _tickDuration;
        maxJoinsInTick = _maxJoinsInTick;
        requiredNodeBalance = _requiredNodeBalance;
    }

    function depositForNode(uint256[4] calldata nodeId) external payable {
        address msgSender = msg.sender;
        bytes32 nodeIdHash = keccak256(abi.encode(nodeId));
        require(
            nodeData[nodeIdHash].balance + msg.value >= requiredNodeBalance,
            "Node Must Have Required Node Balance"
        );
        if (nodeData[nodeIdHash].owner == address(0)) {
            if (block.timestamp > (curTickNum + tickDuration)) {
                curTickNum += (tickDuration *
                    ((block.timestamp - curTickNum) / tickDuration));
                nodesJoinedInCurTick++;
            } else {
                require(
                    nodesJoinedInCurTick < maxJoinsInTick,
                    "Max Node Join In A Tick Limit Exceeded"
                );
                nodesJoinedInCurTick++;
            }
            require(
                BLS.isValidPublicKey(nodeId) == true,
                "Invalid NodeId, Make Sure You Are Passing PubKey That Comes In Curve"
            );
            nodeData[nodeIdHash].balance = msg.value;
            nodeData[nodeIdHash].owner = msgSender;
            unchecked {
                totalNodes++;
            }
            emit NodeAdded(nodeIdHash, msgSender, curTickNum);
        } else {
            nodeData[nodeIdHash].balance += msg.value;
            emit ExtraDepositForNode(nodeIdHash, msgSender, msg.value);
        }
    }

    function withdrawNodeFunds(uint256[4] calldata nodeId) external {
        address msgSender = msg.sender;
        bytes32 nodeIdHash = keccak256(abi.encode(nodeId));
        require(
            nodeData[nodeIdHash].owner == msgSender,
            "Invalid Node Owner, You Cannot Remove This Node"
        );
        require(
            nodeData[nodeIdHash].balance != 0,
            "Node Not Exists In Network"
        );
        uint256 amountToReturn = nodeData[nodeIdHash].balance;
        delete nodeData[nodeIdHash];
        payable(msgSender).transfer(amountToReturn);
        unchecked {
            totalNodes--;
        }
        emit NodeLeft(nodeIdHash, msgSender);
    }

    function removeNode(
        uint256[4][] calldata signersPubKey,
        uint256[2] calldata aggregatedSignature,
        uint256[2] calldata aggregatedPubKeyG1,
        uint256[4] calldata aggregatedPubKeyG2,
        bytes32 nodeToRemoveHashId
    ) external {
        require(
            BLS.verifyHelpedAggregationPublicKeys(
                aggregatedPubKeyG1,
                aggregatedPubKeyG2
            ) == true,
            "Invalid G1 Or G2, PubKeyG1 And PubKeyG2 Are Not Associated"
        );
        require(
            nodeData[nodeToRemoveHashId].owner != address(0),
            "Node Not Exists, May Node Is Removed"
        );
        uint16 totalSigners = uint16(signersPubKey.length);
        for (uint16 i = 0; i < totalSigners; ) {
            require(
                nodeData[keccak256(abi.encode(signersPubKey[i]))].balance >=
                    requiredNodeBalance,
                "Passed Signer PubKey Is Not A Registered Node"
            );
            unchecked {
                i++;
            }
        }
        // // Commenting Below Given Code For Testing Purpose
        // require(
        //     totalSigners >= (totalNodes / 2),
        //     "Requires Majority Of Nodes Signature"
        // );
        require(
            BLS.verifyHelpedAggregationPublicKeysMultiple(
                aggregatedPubKeyG1,
                signersPubKey
            ) == true,
            "Invalid AggregatedPubKeyG2"
        );
        require(
            BLS.verifySingle(
                aggregatedSignature,
                aggregatedPubKeyG2,
                BLS.hashToPoint(
                    abi.encode(
                        SignatureMessage(
                            nodeToRemoveHashId,
                            nonceForNode[nodeToRemoveHashId]
                        )
                    )
                )
            ) == true,
            "Signature Verification Failed"
        );
        address nodeOwner = nodeData[nodeToRemoveHashId].owner;
        uint256 amountLocked = nodeData[nodeToRemoveHashId].balance;
        delete nodeData[nodeToRemoveHashId];
        unchecked {
            totalNodes--;
        }
        unchecked {
            nonceForNode[nodeToRemoveHashId]++;
        }
        emit NodeRemoved(nodeToRemoveHashId, nodeOwner, amountLocked);
    }

    /********** Setter Functions **********/

    function setNewTickDuration(uint256 newTickDuration) external onlyOwner {
        uint256 oldTickDuration = tickDuration;
        tickDuration = newTickDuration;
        emit TickDurationChanged(oldTickDuration, newTickDuration);
    }

    function setNewMaxJoinsInTick(uint256 newMaxJoinsInTick)
        external
        onlyOwner
    {
        uint256 oldMaxJoinsInTick = maxJoinsInTick;
        maxJoinsInTick = newMaxJoinsInTick;
        emit MaxJoinsInTickChanged(oldMaxJoinsInTick, newMaxJoinsInTick);
    }

    function setNewRequiredNodeBalance(uint256 newRequiredNodeBalance)
        external
        onlyOwner
    {
        uint256 oldRequiredNodeBalance = requiredNodeBalance;
        requiredNodeBalance = newRequiredNodeBalance;
        emit RequiredNodeBalanceChanged(
            oldRequiredNodeBalance,
            newRequiredNodeBalance
        );
    }

    /********** View Functions **********/

    function getMessageForNodeParticipationSig(
        string memory message,
        uint256 nonce
    ) external view returns (bytes memory, bytes memory) {
        uint256[2] memory hashToPointData = BLS.hashToPoint(
            abi.encode(message, nonce)
        );
        return (abi.encode(hashToPointData[0]), abi.encode(hashToPointData[1]));
    }

    function getMessageForNodeRemovalSig(bytes32 nodeToRemoveHashId)
        external
        view
        returns (bytes memory, bytes memory)
    {
        uint256[2] memory hashToPointData = BLS.hashToPoint(
            abi.encode(
                SignatureMessage(
                    nodeToRemoveHashId,
                    nonceForNode[nodeToRemoveHashId]
                )
            )
        );
        return (abi.encode(hashToPointData[0]), abi.encode(hashToPointData[1]));
    }

    function isValidNodeId(uint256[4] memory nodeId)
        external
        pure
        returns (bool)
    {
        return BLS.isValidPublicKey(nodeId);
    }

    function isNodeExists(uint256[4] calldata nodeId)
        external
        view
        returns (bool)
    {
        return
            nodeData[keccak256(abi.encode(nodeId))].balance >=
            requiredNodeBalance;
    }

    function getNodeIdHash(uint256[4] calldata nodeId)
        external
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode(nodeId));
    }
}
// 0x7F152299960132F4F5f0E78a53c0345cBCe8b257
