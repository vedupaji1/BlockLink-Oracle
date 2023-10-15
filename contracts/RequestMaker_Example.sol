// SPDX-License-Identifier: LGPL 3.0
pragma solidity 0.8.8;

import "@openzeppelin/contracts/access/Ownable.sol";

// Note: This Contract Specifies The Way Any Requester Contract Have To Make Request And How They They Have To Receive Data.
//       Along With This It Also Show That What Thing Requester Contract Dev Have To Make Sure And How They Have To Design Their Contract.

interface IRequestManager {
    function createRequest() external;
}

contract RequestMaker_Example1 is Ownable {
    event DataReceived(uint256 data);
    event RequestManagerChanged(
        address oldRequestManager,
        address newRequestManager
    );

    IRequestManager public requestManagerIns;

    constructor(address _requestManager) {
        requestManagerIns = IRequestManager(_requestManager);
    }

    receive() external payable {}

    modifier onlyRequestManager() {
        require(
            msg.sender == address(requestManagerIns),
            "Only Request Manager Can Call `handleRequestFulfillment` Method"
        );
        _;
    }

    function makeRequest() external onlyOwner {
        // Add Custom Logic Here
        requestManagerIns.createRequest();
        // Add Custom Logic Here
    }

    function handleRequestFulfillment(
        uint256 data
    ) external onlyRequestManager {
        // Add Custom Logic Here
        emit DataReceived(data);
    }

    function setNewRequestManager(
        address newRequestManager
    ) external onlyOwner {
        address oldRequestManager = address(requestManagerIns);
        requestManagerIns = IRequestManager(newRequestManager);
        emit RequestManagerChanged(oldRequestManager, newRequestManager);
    }
}
// 0xAFb1ce044032c6A7CCBBcaaebab54820398907A7