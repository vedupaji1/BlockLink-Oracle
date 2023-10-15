// SPDX-License-Identifier: LGPL 3.0
pragma solidity 0.8.8;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./BLS.sol";

// Note: This Contract Is For BlockLink Oracle, It Is Under Development And Not Optimized

interface INodeManager {
    function isNodeExists(uint256[4] calldata nodeId)
        external
        view
        returns (bool);

    function totalNodes() external view returns (uint256);
}

contract RequestManager is Ownable {
    enum RequestStatus {
        NotCreated,
        Created,
        Cancelled,
        Fulfilled
    }

    struct Request {
        uint256 createdAt;
        RequestStatus status;
    }

    struct RequestTopic {
        address creator;
        uint256 balance;
        uint256 pendingRequests;
        uint256 totalRequests;
        uint32 maxGasLimit;
        mapping(uint256 => Request) requestsData;
    }

    struct SignatureMessage {
        uint256 chainId;
        address verifier;
        address recipient;
        uint256 requestNum;
        uint256 data;
        uint32 maxGasLimit;
    }

    event RequestTopicCreated(
        address indexed requestBeneficiary,
        address indexed creator
    );
    event DepositedForRequestTopic(address indexed depositor, uint256 amount);
    event WithdrawRequestTopicFunds(uint256 amount);
    event RequestCreated(address indexed receiver, uint256 indexed requestNum);
    event RequestCanceled(address indexed receiver, uint256 indexed requestNum);
    event RequestFulFilled(
        address indexed receiver,
        uint256 indexed requestNum,
        uint256 indexed data,
        uint256 requestTotalCost
    );
    event ServiceCostChanged(
        uint256 indexed oldServiceCost,
        uint256 indexed newServiceCost
    );
    event AmountToLockForEachRequestChanged(
        uint256 oldAmountToLockForEachRequest,
        uint256 newAmountToLockForEachRequest
    );
    event RequestTimeoutChanged(
        uint32 oldRequestTimeout,
        uint32 newRequestTimeout
    );
    event DefaultMaxGasLimitChanged(
        uint32 oldDefaultMaxGasLimit,
        uint32 newDefaultMaxGasLimit
    );
    event NodeManagerChanged(
        address indexed oldNodeManager,
        address indexed newNodeManager
    );
    event TopicRequestMaxGasLimitChanged(
        address indexed oldMaxGasLimit,
        address indexed newMaxGasLimit
    );

    uint256 public serviceCost;
    uint256 public amountToLockForEachRequest;
    uint32 public requestTimeout;
    uint32 public defaultMaxGasLimit;
    INodeManager public nodeManagerIns;
    mapping(address => RequestTopic) private requestTopicData;
    mapping(uint256 => bool) public isUsedData;

    constructor(
        uint256 _serviceCost,
        uint256 _amountToLockForEachRequest,
        uint32 _requestTimeout,
        uint32 _defaultMaxGasLimit,
        address _nodeManager
    ) {
        serviceCost = _serviceCost;
        amountToLockForEachRequest = _amountToLockForEachRequest;
        requestTimeout = _requestTimeout;
        defaultMaxGasLimit = _defaultMaxGasLimit;
        nodeManagerIns = INodeManager(_nodeManager);
    }

    function createRequestTopic(address requestBeneficiary, uint32 maxGasLimit)
        external
        payable
    {
        address msgSender = msg.sender;
        require(
            msg.value > amountToLockForEachRequest,
            "Initial Balance Must Be Significant For Single Request"
        );
        require(
            maxGasLimit <= 30000000,
            "`maxGasLimit` Must Be Lesser Than 30M "
        );
        require(
            maxGasLimit >= defaultMaxGasLimit,
            "`maxGasLimit` Must Be Greater Or Equal To `defaultMaxGasLimit`"
        );
        RequestTopic storage requestTopic = requestTopicData[
            requestBeneficiary
        ];
        require(
            requestTopic.creator == address(0),
            "Request Topic Is Already Created"
        );
        requestTopic.creator = msgSender;
        requestTopic.balance = msg.value;
        requestTopic.maxGasLimit = maxGasLimit;
        emit RequestTopicCreated(requestBeneficiary, msgSender);
    }

    function depositForRequestTopic(address requestBeneficiary)
        external
        payable
    {
        require(
            requestTopicData[requestBeneficiary].creator != address(0),
            "Request Topic Not Exists"
        );
        requestTopicData[requestBeneficiary].balance += msg.value;
        emit DepositedForRequestTopic(msg.sender, msg.value);
    }

    function withdrawRequestTopicFunds(
        address requestBeneficiary,
        uint256 amount
    ) external payable {
        RequestTopic storage requestTopic = requestTopicData[
            requestBeneficiary
        ];
        require(
            msg.sender == requestTopic.creator,
            "Only RequestTopic Creator Can Withdraw Remaining Funds"
        );
        require(
            amount <=
                (requestTopic.balance -
                    (requestTopic.pendingRequests *
                        amountToLockForEachRequest)),
            "Insufficient Amount Available To Withdraw"
        );
        requestTopic.balance -= amount;
        payable(requestTopic.creator).transfer(amount);
        emit WithdrawRequestTopicFunds(amount);
    }

    function createRequest() external {
        address msgSender = msg.sender;
        require(
            msgSender.code.length > 0,
            "Only Smart Contract Can Create Request"
        );
        RequestTopic storage requestTopic = requestTopicData[msgSender];
        require(
            (requestTopic.balance -
                (requestTopic.pendingRequests * amountToLockForEachRequest)) >=
                amountToLockForEachRequest,
            "Insufficient Funds For Creating Request"
        );
        uint256 curRequestNum = ++requestTopic.totalRequests;
        requestTopic.pendingRequests++;
        requestTopic.requestsData[curRequestNum] = Request({
            createdAt: block.timestamp,
            status: RequestStatus.Created
        });

        emit RequestCreated(msgSender, curRequestNum);
    }

    function cancelRequest(address requestBeneficiary, uint256 requestNum)
        external
    {
        RequestTopic storage requestTopic = requestTopicData[
            requestBeneficiary
        ];
        require(
            msg.sender == requestTopic.creator,
            "Only RequestTopic Creator Can Withdraw Remaining Funds"
        );
        require(
            requestTopic.requestsData[requestNum].status ==
                RequestStatus.Created,
            "Request Status Must Be `Created`"
        );
        require(
            (block.timestamp -
                requestTopic.requestsData[requestNum].createdAt) >
                requestTimeout,
            "Request Can Only Be Cancelled After Request Timeout"
        );
        requestTopic.requestsData[requestNum].status = RequestStatus.Cancelled;
        requestTopic.pendingRequests--;
        emit RequestCanceled(requestBeneficiary, requestNum);
    }

    function fullFillRequest(
        uint256[4][] calldata signersPubKey,
        uint256[2] calldata aggregatedSignature,
        uint256[2] calldata aggregatedPubKeyG1,
        uint256[4] calldata aggregatedPubKeyG2,
        address requestBeneficiary,
        uint256 requestNum,
        uint256 data
    ) external {
        RequestTopic storage requestTopic = requestTopicData[
            requestBeneficiary
        ];
        require(
            requestTopic.requestsData[requestNum].status ==
                RequestStatus.Created,
            "Request Status Must Be Created`"
        );
        require(
            isUsedData[data] == false,
            "Passed Data Is Already Used, Passed Data Must Be All Time Unique"
        );
        require(
            BLS.verifyHelpedAggregationPublicKeys(
                aggregatedPubKeyG1,
                aggregatedPubKeyG2
            ) == true,
            "Invalid G1 Or G2, PubKeyG1 And PubKeyG2 Are Not Associated"
        );
        uint16 totalSigners = uint16(signersPubKey.length);
        for (uint16 i = 0; i < totalSigners; ) {
            require(
                nodeManagerIns.isNodeExists(signersPubKey[i]) == true,
                "Passed Signer PubKey Is Not A Registered Node"
            );
            unchecked {
                i++;
            }
        }
        // // Commenting Below Given Code For Testing Purpose
        // require(
        //     totalSigners >= (nodeManagerIns.totalNodes() / 2),
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
                            block.chainid,
                            address(this),
                            requestBeneficiary,
                            requestNum,
                            data,
                            requestTopic.maxGasLimit
                        )
                    )
                )
            ) == true,
            "Signature Verification Failed"
        );

        (bool success, ) = requestBeneficiary.call(
            abi.encodeWithSignature("handleRequestFulfillment(uint256)", data)
        );
        require(
            success == true,
            "Failed To Successfully Deliver Data To Receiver, Something Wrong From Receiver Side"
        );

        isUsedData[data] = true;
        requestTopic.requestsData[requestNum].status = RequestStatus.Fulfilled;
        requestTopic.pendingRequests--;
        uint256 actualGasUsed = (requestTopic.maxGasLimit - gasleft()) + 9891; // 9891 (5091 + 4800) Is Amount Of Required To execute Below Given Code
        uint256 requestTotalCost = (actualGasUsed * tx.gasprice) + serviceCost;
        requestTopic.balance -= requestTotalCost;
        emit RequestFulFilled(
            requestBeneficiary,
            requestNum,
            data,
            actualGasUsed
        );
    }

    /********** Setter Functions **********/

    function setNewServiceCost(uint256 newServiceCost) external onlyOwner {
        uint256 oldServiceCost = serviceCost;
        serviceCost = newServiceCost;
        emit ServiceCostChanged(oldServiceCost, newServiceCost);
    }

    function setNewAmountToLockForEachRequest(
        uint256 newAmountToLockForEachRequest
    ) external onlyOwner {
        uint256 oldAmountToLockForEachRequest = amountToLockForEachRequest;
        amountToLockForEachRequest = newAmountToLockForEachRequest;
        emit AmountToLockForEachRequestChanged(
            oldAmountToLockForEachRequest,
            newAmountToLockForEachRequest
        );
    }

    function seNewtRequestTimeout(uint32 newRequestTimeout) external onlyOwner {
        uint32 oldRequestTimeout = requestTimeout;
        requestTimeout = newRequestTimeout;
        emit RequestTimeoutChanged(oldRequestTimeout, newRequestTimeout);
    }

    function setNewDefaultMaxGasLimit(uint32 newDefaultMaxGasLimit)
        external
        onlyOwner
    {
        require(
            newDefaultMaxGasLimit <= 30000000,
            "`newDefaultMaxGasLimit` Must Be Lesser Than 30M "
        );
        uint32 oldDefaultMaxGasLimit = defaultMaxGasLimit;
        defaultMaxGasLimit = newDefaultMaxGasLimit;
        emit DefaultMaxGasLimitChanged(
            oldDefaultMaxGasLimit,
            newDefaultMaxGasLimit
        );
    }

    function setTopicRequestNewMaxGasLimit(uint32 newMaxGasLimit)
        external
        onlyOwner
    {
        require(
            newMaxGasLimit <= 30000000,
            "`newMaxGasLimit` Must Be Lesser Than 30M "
        );
        address msgSender = msg.sender;
        require(
            msgSender == requestTopicData[msgSender].creator,
            "Only Request Topic Creator Can Access This Method"
        );
        uint32 oldMaxGasLimit = requestTopicData[msgSender].maxGasLimit;
        requestTopicData[msgSender].maxGasLimit = newMaxGasLimit;
        emit DefaultMaxGasLimitChanged(oldMaxGasLimit, newMaxGasLimit);
    }

    function setNewNodeManager(address newNodeManager) external onlyOwner {
        address oldNodeManager = address(nodeManagerIns);
        nodeManagerIns = INodeManager(newNodeManager);
        emit NodeManagerChanged(oldNodeManager, newNodeManager);
    }

    /********** View Functions **********/

    function isValidPublicKey(uint256[4] memory publicKey)
        external
        pure
        returns (bool)
    {
        return BLS.isValidPublicKey(publicKey);
    }

    function isValidSignature(uint256[2] memory signature)
        external
        pure
        returns (bool)
    {
        return BLS.isValidSignature(signature);
    }

    function getMessageForRequestFulFillmentSig(
        address recipient,
        uint256 requestNum,
        uint256 data
    ) external view returns (bytes memory, bytes memory) {
        uint256[2] memory hashToPointData = BLS.hashToPoint(
            abi.encode(
                SignatureMessage(
                    block.chainid,
                    address(this),
                    recipient,
                    requestNum,
                    data,
                    requestTopicData[recipient].maxGasLimit
                )
            )
        );
        return (abi.encode(hashToPointData[0]), abi.encode(hashToPointData[1]));
    }

    function hashToPoint(bytes memory data)
        external
        view
        returns (bytes memory, bytes memory)
    {
        uint256[2] memory hashToPointData = BLS.hashToPoint(data);
        return (abi.encode(hashToPointData[0]), abi.encode(hashToPointData[1]));
    }

    function verifyHelpedAggregationPublicKeys(
        uint256[2] memory pubkeyG1,
        uint256[4] memory pubkeyG2
    ) external view returns (bool) {
        return BLS.verifyHelpedAggregationPublicKeys(pubkeyG1, pubkeyG2);
    }

    function verifyHelpedAggregationPublicKeysMultiple(
        uint256[2] memory aggPubkeyG1,
        uint256[4][] memory pubkeysG2
    ) external view returns (bool) {
        return
            BLS.verifyHelpedAggregationPublicKeysMultiple(
                aggPubkeyG1,
                pubkeysG2
            );
    }

    function requestTopicAvailableFunds(address requestBeneficiary)
        external
        view
        returns (uint256)
    {
        return
            requestTopicData[requestBeneficiary].balance -
            (requestTopicData[requestBeneficiary].pendingRequests *
                amountToLockForEachRequest);
    }

    function getRequestTopicData(address requestBeneficiary)
        external
        view
        returns (
            address,
            uint256,
            uint256,
            uint256,
            uint32
        )
    {
        return (
            requestTopicData[requestBeneficiary].creator,
            requestTopicData[requestBeneficiary].balance,
            requestTopicData[requestBeneficiary].pendingRequests,
            requestTopicData[requestBeneficiary].totalRequests,
            requestTopicData[requestBeneficiary].maxGasLimit
        );
    }

    function getRequestData(address requestBeneficiary, uint256 requestNum)
        external
        view
        returns (uint256, RequestStatus)
    {
        return (
            requestTopicData[requestBeneficiary]
                .requestsData[requestNum]
                .createdAt,
            requestTopicData[requestBeneficiary].requestsData[requestNum].status
        );
    }
}
// 0x7Ae4f1C3371e1c882c5e653454EEeab5f22dAabC
