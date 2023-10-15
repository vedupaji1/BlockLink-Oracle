package temp11

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/http"
	"strings"
	constants "temp11/constants"
	"temp11/nodeManager"
	"temp11/requestManager"

	"github.com/ethereum/go-ethereum/common"
	logger "github.com/inconshreveable/log15"
	// "github.com/syndtr/goleveldb/leveldb"
)

type PeerSignatureMessage struct {
	PlatformName string
	Nonce        int
	Data         interface{}
}

type RegistrationRequestSignatureMessage struct {
	PlatformName     string
	RequestMessage   string
	PeerAddrs        PeerReqAndListenAddr
	PeerRegisteredId PeersRegisteredId
	Nonce            int
}

type EmptyArgs struct{}

type EmptyReply struct{}

type PeerDataReply struct {
	// db                   *leveldb.DB
	PeersList                 []string
	PeerIndex                 map[common.Hash]int
	PeersAddr                 map[string]PeerReqAndListenAddr
	PeersRegisteredId         map[string]PeersRegisteredId
	PeersUsedNonces           map[string]map[int]bool
	LeaderPeer                string
	LeaderPeerIndex           int
	NextLeaderChangeTime      int64
	DataForRequest            map[string]*DataRequestInfo
	PeersParticipationRequest map[string]int
}

type RegisterNewPeerArgs struct {
	NodeId                 string
	Signature              string
	RegisteredKeySignature [3]*big.Int
	Nonce                  int
	RequestMessage         string
	RegisteredKeyHash      common.Hash
	PeerRegisteredId       PeersRegisteredId
	RequesterAddrs         PeerReqAndListenAddr
}

type ReceiveRegisterNewPeerArgs struct {
	LeaderSignature        string
	Nonce                  int
	RequesterSignatureInfo RegisterNewPeerArgs
}

type StoreDataArgs struct {
	Id   int
	Data string
}

type ReceiveDataArgs struct {
	LeaderSignature string
	Nonce           int
	DataToStore     StoreDataArgs
}

type RandDataRequestSigMessage struct {
	TxHash   string
	Receiver common.Address
}

type UnFinalizedRandData struct {
	RandData string
	TxHash   string
}

type FinalizedRandData struct {
	RandData string
	TxHash   string
}

type RandDataForRequestArgs struct {
	LeaderSignature string
	Nonce           int
	TxHash          string
	Receiver        common.Address
	RequestNum      *big.Int
}

type RandDataForRequestReply struct {
	UnFinalizedRandData string
	Signature           string
	Nonce               int
}

type ReceiveFinalizedRandDataArgs struct {
	FinalizedRandData             string
	LeaderSignature               string
	Nonce                         int
	TxHash                        string
	SupportForUnFinalizedRandData map[string]UnFinalizedRandDataSupportInfo
}

type ReceiveFinalizedRandDataReply struct {
	SignatureForETH string
}

type ReceiveFinalizedRandDataSupportInfoArgs struct {
	LeaderSignature             string
	Nonce                       int
	TxHash                      string
	SupportForFinalizedRandData map[string]FinalizedRandDataSupportInfo
}

type ReceiveFulFillRequestTxHashArgs struct {
	LeaderSignature    string
	Nonce              int
	RequestTxHash      string
	FulFillRequestHash string
}

type RequestAndFulFullTxHash struct {
	RequestTxHash      string
	FulFillRequestHash string
}

type AgreeToRemovePeerFromNodeManagerArgs struct {
	LeaderSignature    string
	Nonce              int
	NodeToRemoveHashId common.Hash
}

type AgreeToRemovePeerFromNodeManagerReply struct {
	SignatureToRemovePeer string
}

func (node *Node) GetNodeId(r *http.Request, req *EmptyArgs, res *string) error {
	logger.Info("Received `GetNodeId` Request", "RemoteAddr", r.RemoteAddr)
	*res = node.Id
	return nil
}

func (node *Node) GetPeersData(r *http.Request, req *EmptyArgs, res *PeerDataReply) error {
	logger.Info("Received `GetPeersData` Request", "RemoteAddr", r.RemoteAddr)
	res.PeersList = node.PeerList
	res.PeerIndex = node.PeerIndex
	res.PeersAddr = node.PeersAddr
	res.PeersRegisteredId = node.PeersRegisteredId
	res.LeaderPeer = node.PeerList[node.LeaderPeerIndex]
	res.LeaderPeerIndex = node.LeaderPeerIndex
	res.NextLeaderChangeTime = node.NextLeaderChangeTime
	res.PeersUsedNonces = node.PeersUsedNonces
	res.DataForRequest = node.DataForRequest
	res.PeersParticipationRequest = SyncMapToStringIntMap(node.PeersParticipationRequest)
	return nil
}

func (node *Node) RegisterNewPeer(r *http.Request, req *RegisterNewPeerArgs, res *PeerDataReply) error {
	logger.Info("Received `RegisterNewPeer` Request", "RemoteAddr", r.RemoteAddr)
	if node.PeerList[node.LeaderPeerIndex] != node.Id {
		logger.Error("Received Invalid Data Storage Request", "RemoteAddr", r.RemoteAddr)
		return fmt.Errorf("only leader Node can store data")
	}
	if _, ok := node.PeersAddr[req.NodeId]; ok {
		logger.Error("Invalid Registration Request, Peer Is Already Registered", "PeerId", req.NodeId, "RemoteAddr", r.RemoteAddr)
		return fmt.Errorf("peer is already registered")
	}

	logger.Info("Pinging Registration Requester Peer", "PeerListenAddr", req.RequesterAddrs.ListenAddr)
	resp, err := http.Get(req.RequesterAddrs.ListenAddr)
	if err != nil {
		logger.Error("Couldn't Connecting To Registration Requester Peer", "PeerListenAddr", req.RequesterAddrs.ListenAddr, "Error", err)
		return fmt.Errorf("couldn't connecting to registration requester peer: %v", err)
	}
	if resp.StatusCode == http.StatusOK {
		logger.Info("Pinged Successfully To Registration Requester Peer", "ListenAddr", req.RequesterAddrs.ListenAddr)
		resp.Body.Close()
	} else {
		logger.Error("Bad Request", "Response", resp)
		return fmt.Errorf("bad request: %v", resp)
	}

	requesterRegisteredIdHash, err := HashRegisteredId(node.bn128_bls.ParsePubKey(req.PeerRegisteredId.G2))
	if err != nil {
		return err
	}
	if _, ok := node.PeersParticipationRequest.Load(requesterRegisteredIdHash); !ok {
		return fmt.Errorf("requester have not done registration on ethereum")
	}

	if err := node.verifySignature(req.NodeId, req.Signature, req.Nonce, RegistrationRequestSignatureMessage{
		PlatformName:     constants.PlatformName,
		RequestMessage:   req.RequestMessage,
		PeerAddrs:        req.RequesterAddrs,
		PeerRegisteredId: req.PeerRegisteredId,
		Nonce:            req.Nonce,
	}); err != nil {
		logger.Info("Requester Signature Verification Failed")
		return err
	}
	logger.Info("Requester Signature Verification Done", "Signature", req.Signature)

	contract, err := nodeManager.NewNodeManager(common.HexToAddress(constants.NodeManagerContract), node.ethClientRPC)
	if err != nil {
		logger.Error("Failed To Create Request Manager Contract Instance", "Error", err)
		return fmt.Errorf("failed to create request manager contract instance")
	}
	messageX, messageY, err := contract.GetMessageForNodeParticipationSig(nil, constants.RegistrationRequestMessage, big.NewInt(int64(req.Nonce)))
	if err != nil {
		logger.Error("Failed To Get Signature Message", "Error", err)
		log.Panic("Failed To Get Signature Message")
	}
	if ok, err := node.bn128_bls.VerifySignature(req.RegisteredKeySignature, req.PeerRegisteredId.G2, hex.EncodeToString(messageX), hex.EncodeToString(messageY)); !ok {
		logger.Error("Failed To Verify RegisteredId Signature", "Error", err)
		return fmt.Errorf("failed to verify registeredId signature: %v", err)
	}
	logger.Info("Requester RegisteredKey Signature Is Verified", "RegisteredKeySignature", req.RegisteredKeySignature)
	node.PeersParticipationRequest.Delete(requesterRegisteredIdHash)

	leaderSignatureMessage, nonce, err := node.generatePeerSignatureMessage(req.Signature)
	if err != nil {
		logger.Error("Failed To Generate Signature Message")
		return fmt.Errorf("something went wrong while signature message generation: %v", err)
	}
	leaderPeerSignature := node.nodeKey.SignByte(leaderSignatureMessage).SerializeToHexStr()
	logger.Info("Signature Generated By Leader Node", "Signature", leaderPeerSignature)
	resData := make(map[string]interface{})
	node.broadCastMessage(CopyStringArray(node.PeerList), "Node.ReceiveRegisterNewPeer", ReceiveRegisterNewPeerArgs{
		LeaderSignature:        leaderPeerSignature,
		Nonce:                  nonce,
		RequesterSignatureInfo: *req,
	}, resData)
	if err != nil {
		logger.Error("Failed To Generate Signature Message")
		return fmt.Errorf("something went wrong while signature message generation: %v", err)
	}
	peerRegisteredIdHash, err := HashRegisteredId(node.bn128_bls.ParsePubKey(req.PeerRegisteredId.G2))
	if err != nil {
		logger.Error("", "Error", err)
		return err
	}
	node.PeersAddr[req.NodeId] = req.RequesterAddrs
	node.PeerIndex[peerRegisteredIdHash] = len(node.PeerList)
	node.PeerList = append(node.PeerList, req.NodeId)
	node.PeersRegisteredId[req.NodeId] = req.PeerRegisteredId

	res.PeersAddr = node.PeersAddr
	res.PeersList = node.PeerList
	res.PeerIndex = node.PeerIndex
	res.PeersRegisteredId = node.PeersRegisteredId
	res.LeaderPeer = node.PeerList[node.LeaderPeerIndex]
	res.LeaderPeerIndex = node.LeaderPeerIndex
	res.NextLeaderChangeTime = node.NextLeaderChangeTime
	res.PeersUsedNonces = node.PeersUsedNonces
	res.DataForRequest = node.DataForRequest
	res.PeersParticipationRequest = SyncMapToStringIntMap(node.PeersParticipationRequest)
	logger.Info("New Peer Added", "PeerId", req.NodeId)
	return nil
}

func (node *Node) ReceiveRegisterNewPeer(r *http.Request, req *ReceiveRegisterNewPeerArgs, res *EmptyReply) error {
	logger.Info("Received `ReceiveRegisterNewPeer` Request", "RemoteAddr", r.RemoteAddr)
	leaderPeerIp, _, _ := net.SplitHostPort(strings.TrimPrefix(node.PeersAddr[node.PeerList[node.LeaderPeerIndex]].RequestAddr, "http://"))
	requesterIp, _, _ := net.SplitHostPort(r.RemoteAddr)
	if requesterIp != leaderPeerIp && requesterIp != "127.0.0.1" {
		logger.Error("Received Invalid Request", "RemoteAddr", r.RemoteAddr)
		return fmt.Errorf("only leader node can send request for storing data")
	}
	logger.Info("Received Data From Node", "Data", *req)

	leaderPeerID := node.PeerList[node.LeaderPeerIndex]
	if err := node.verifySignature(leaderPeerID, req.LeaderSignature, req.Nonce, PeerSignatureMessage{
		PlatformName: constants.PlatformName,
		Nonce:        req.Nonce,
		Data:         req.RequesterSignatureInfo.Signature,
	}); err != nil {
		logger.Error("Leader Signature Verification Failed")
		return err
	}

	// Not Verifying Finalized RandData, Because Of Not Proper Implementation Of BFT
	logger.Info("Leader Signature Verification Done", "Signature", req.LeaderSignature)
	if err := node.verifySignature(req.RequesterSignatureInfo.NodeId, req.RequesterSignatureInfo.Signature, req.RequesterSignatureInfo.Nonce, RegistrationRequestSignatureMessage{
		PlatformName:     constants.PlatformName,
		RequestMessage:   req.RequesterSignatureInfo.RequestMessage,
		PeerAddrs:        req.RequesterSignatureInfo.RequesterAddrs,
		PeerRegisteredId: req.RequesterSignatureInfo.PeerRegisteredId,
		Nonce:            req.RequesterSignatureInfo.Nonce,
	}); err != nil {
		logger.Info("Requester Signature Verification Failed")
		return err
	}
	logger.Info("Requester Signature Verification Done", "Signature", req.RequesterSignatureInfo.Signature)

	requesterRegisteredIdHash, err := HashRegisteredId(node.bn128_bls.ParsePubKey(req.RequesterSignatureInfo.PeerRegisteredId.G2))
	if err != nil {
		return err
	}
	if _, ok := node.PeersParticipationRequest.Load(requesterRegisteredIdHash); !ok {
		return fmt.Errorf("requester have not done registration on ethereum")
	}
	contract, err := nodeManager.NewNodeManager(common.HexToAddress(constants.NodeManagerContract), node.ethClientRPC)
	if err != nil {
		logger.Error("Failed To Create Request Manager Contract Instance", "Error", err)
		return fmt.Errorf("failed to create request manager contract instance")
	}
	messageX, messageY, err := contract.GetMessageForNodeParticipationSig(nil, constants.RegistrationRequestMessage, big.NewInt(int64(req.RequesterSignatureInfo.Nonce)))
	if err != nil {
		logger.Error("Failed To Get Signature Message", "Error", err)
		log.Panic("Failed To Get Signature Message")
	}
	if ok, err := node.bn128_bls.VerifySignature(req.RequesterSignatureInfo.RegisteredKeySignature, req.RequesterSignatureInfo.PeerRegisteredId.G2, hex.EncodeToString(messageX), hex.EncodeToString(messageY)); !ok {
		logger.Error("Failed To Verify RegisteredKey Signature", "Error", err)
		return fmt.Errorf("failed to verify registeredKey signature: %v", err)
	}
	logger.Info("Requester RegisteredKey Signature Is Verified", "RegisteredKeySignature", req.RequesterSignatureInfo.RegisteredKeySignature)

	peerRegisteredIdHash, err := HashRegisteredId(node.bn128_bls.ParsePubKey(req.RequesterSignatureInfo.PeerRegisteredId.G2))
	if err != nil {
		logger.Error("", "Error", err)
		return err
	}
	node.PeersParticipationRequest.Delete(requesterRegisteredIdHash)
	node.PeersAddr[req.RequesterSignatureInfo.NodeId] = req.RequesterSignatureInfo.RequesterAddrs
	node.PeerIndex[peerRegisteredIdHash] = len(node.PeerList)
	node.PeerList = append(node.PeerList, req.RequesterSignatureInfo.NodeId)
	node.PeersRegisteredId[req.RequesterSignatureInfo.NodeId] = req.RequesterSignatureInfo.PeerRegisteredId
	logger.Info("New Peer Added", "PeerId", req.RequesterSignatureInfo.NodeId)
	return nil
}

func (node *Node) StoreData(r *http.Request, req *StoreDataArgs, res *EmptyReply) error {
	logger.Info("Received `StoreData` Request", "RemoteAddr", r.RemoteAddr)
	if node.PeerList[node.LeaderPeerIndex] != node.Id {
		logger.Error("Received Invalid Data Storage Request", "RemoteAddr", r.RemoteAddr)
		return fmt.Errorf("only leader node can store data")
	}
	signatureMessage, nonce, err := node.generatePeerSignatureMessage(*req)
	if err != nil {
		logger.Error("Failed To Generate Signature Message")
		return fmt.Errorf("something went wrong while signature message generation: %v", err)
	}
	signature := node.nodeKey.SignByte(signatureMessage).SerializeToHexStr()
	logger.Info("Signature Generated By Leader Node", "Signature", signature)
	resData := make(map[string]interface{})
	node.broadCastMessage(CopyStringArray(node.PeerList), "Node.ReceiveDataToStore", ReceiveDataArgs{
		LeaderSignature: signature,
		Nonce:           nonce,
		DataToStore:     *req,
	}, resData)
	return nil
}

func (node *Node) ReceiveDataToStore(r *http.Request, req *ReceiveDataArgs, res *EmptyReply) error {
	logger.Info("Received `ReceiveDataToStore` Request", "RemoteAddr", r.RemoteAddr)
	leaderPeerIp, _, _ := net.SplitHostPort(strings.TrimPrefix(node.PeersAddr[node.PeerList[node.LeaderPeerIndex]].RequestAddr, "http://"))
	requesterIp, _, _ := net.SplitHostPort(r.RemoteAddr)
	if requesterIp != leaderPeerIp && requesterIp != "127.0.0.1" {
		logger.Error("Received Invalid Request", "RemoteAddr", r.RemoteAddr)
		return fmt.Errorf("only leader node can send request for storing data")
	}
	logger.Info("Received Data From Node", "Data", *req)

	leaderPeerID := node.PeerList[node.LeaderPeerIndex]
	if err := node.verifySignature(leaderPeerID, req.LeaderSignature, req.Nonce, PeerSignatureMessage{
		PlatformName: constants.PlatformName,
		Nonce:        req.Nonce,
		Data:         req.DataToStore,
	}); err != nil {
		return err
	}
	logger.Info("Dummy: Node Data Is Stored")
	return nil
}

func (node *Node) GetRandDataForRequest(r *http.Request, req *RandDataForRequestArgs, res *RandDataForRequestReply) error {
	logger.Info("Received `GetRandDataForRequest` Request", "RemoteAddr", r.RemoteAddr)
	leaderPeerIp, _, _ := net.SplitHostPort(strings.TrimPrefix(node.PeersAddr[node.PeerList[node.LeaderPeerIndex]].RequestAddr, "http://"))
	requesterIp, _, _ := net.SplitHostPort(r.RemoteAddr)
	if requesterIp != leaderPeerIp && requesterIp != "127.0.0.1" {
		logger.Error("Received Invalid Request", "RemoteAddr", r.RemoteAddr)
		return fmt.Errorf("only leader node can send request for storing data")
	}
	logger.Info("Received Data From Node", "Data", *req)

	leaderPeerID := node.PeerList[node.LeaderPeerIndex]
	if err := node.verifySignature(leaderPeerID, req.LeaderSignature, req.Nonce, PeerSignatureMessage{
		PlatformName: constants.PlatformName,
		Nonce:        req.Nonce,
		Data: RandDataRequestSigMessage{
			TxHash:   req.TxHash,
			Receiver: req.Receiver,
		},
	}); err != nil {
		return err
	}

	randNum, err := GenBigIntRand()
	if err != nil {
		logger.Error("Failed To Generate Random Number")
		return err
	}
	logger.Info("Random Number For Request Is Generated", "RandNumber", randNum)
	signatureMessage, nonce, err := node.generatePeerSignatureMessage(
		UnFinalizedRandData{
			RandData: randNum.String(),
			TxHash:   req.TxHash,
		})
	if err != nil {
		logger.Error("Failed To Generate Signature Message")
		return fmt.Errorf("something went wrong while signature message generation: %v", err)
	}
	signature := node.nodeKey.SignByte(signatureMessage).SerializeToHexStr()
	logger.Info("Signature Generated For UnFinalized Request", "Signature", signature)

	node.mutexForDataRequest.Lock()
	node.DataForRequest[req.TxHash] = &DataRequestInfo{
		FinalizedRandData:             big.NewInt(0),
		LeaderNodeId:                  node.PeerList[node.LeaderPeerIndex],
		SupportForUnFinalizedRandData: make(map[string]UnFinalizedRandDataSupportInfo),
		SupportForFinalizedRandData:   make(map[string]FinalizedRandDataSupportInfo),
		Receiver:                      req.Receiver,
		RequestNum:                    req.RequestNum,
	}
	node.DataForRequest[req.TxHash].SupportForUnFinalizedRandData[node.Id] = UnFinalizedRandDataSupportInfo{
		RandData:  randNum,
		Signature: signature,
		Nonce:     nonce,
	}
	node.mutexForDataRequest.Unlock()
	res.UnFinalizedRandData = randNum.String()
	res.Signature = signature
	res.Nonce = nonce
	logger.Info("UnFinalized RandData Is Shared With Leader Peer", "ResData", res)

	return nil
}

func (node *Node) ReceiveFinalizedRandData(r *http.Request, req *ReceiveFinalizedRandDataArgs, res *ReceiveFinalizedRandDataReply) error {
	logger.Info("Received `ReceiveFinalizedRandData` Request", "RemoteAddr", r.RemoteAddr)
	leaderPeerIp, _, _ := net.SplitHostPort(strings.TrimPrefix(node.PeersAddr[node.PeerList[node.LeaderPeerIndex]].RequestAddr, "http://"))
	requesterIp, _, _ := net.SplitHostPort(r.RemoteAddr)
	if requesterIp != leaderPeerIp && requesterIp != "127.0.0.1" {
		logger.Error("Received Invalid Request", "RemoteAddr", r.RemoteAddr)
		return fmt.Errorf("only leader node can send request for storing data")
	}
	logger.Info("Received Data From Node", "Data", *req)

	leaderPeerID := node.PeerList[node.LeaderPeerIndex]
	if err := node.verifySignature(leaderPeerID, req.LeaderSignature, req.Nonce, PeerSignatureMessage{
		PlatformName: constants.PlatformName,
		Nonce:        req.Nonce,
		Data: FinalizedRandData{
			RandData: req.FinalizedRandData,
			TxHash:   req.TxHash,
		},
	}); err != nil {
		return err
	}

	// Not Verifying Finalized RandData, Because Of Not Proper Implementation Of BFT
	for peerId, data := range req.SupportForUnFinalizedRandData {
		if node.Id != peerId {
			if err := node.verifySignature(peerId, data.Signature, data.Nonce, PeerSignatureMessage{
				PlatformName: constants.PlatformName,
				Nonce:        data.Nonce,
				Data: UnFinalizedRandData{
					RandData: data.RandData.String(),
					TxHash:   req.TxHash,
				},
			}); err != nil {
				return err
			}
		}
	}

	contract, err := requestManager.NewRequestManager(common.HexToAddress(constants.RequestMangerContract), node.ethClientRPC)
	if err != nil {
		logger.Error("Failed To Create Request Fulfiller Contract Instance", "Error", err)
		return fmt.Errorf("something went wrong while signature message generation: %v", err)
	}
	finalizedRandData, ok := new(big.Int).SetString(req.FinalizedRandData, 10)
	if !ok {
		logger.Error("Failed To Parse FinalizedRandData")
		return fmt.Errorf("failed to parse finalizedRandData")
	}
	node.mutexForDataRequest.RLock()
	requestReceiver := node.DataForRequest[req.TxHash].Receiver
	requestNum := node.DataForRequest[req.TxHash].RequestNum
	node.mutexForDataRequest.RUnlock()
	messageX, messageY, err := contract.GetMessageForRequestFulFillmentSig(nil, requestReceiver, requestNum, finalizedRandData)
	if err != nil {
		logger.Error("Failed To Get Signature Message", "Error", err)
		return fmt.Errorf("something went wrong while getting signature message: %v", err)
	}
	logger.Info("Signature Message For `bn128_bls` Signature Generated", "MessageX", messageX, "MessageY", messageY)

	signatureForETH, err := node.bn128_bls.GenerateSignature(node.nodeRegisteredIdKeyPair, hex.EncodeToString(messageX), hex.EncodeToString(messageY))
	if err != nil {
		logger.Error("Failed To Generate Signature For Ethereum")
		return fmt.Errorf("something went wrong while generating signature for ethereum: %v", err)
	}

	node.mutexForDataRequest.Lock()
	node.DataForRequest[req.TxHash].FinalizedRandData = finalizedRandData
	node.DataForRequest[req.TxHash].LeaderNodeSignature = req.LeaderSignature
	node.DataForRequest[req.TxHash].SignatureNonce = req.Nonce
	node.DataForRequest[req.TxHash].SupportForUnFinalizedRandData = req.SupportForUnFinalizedRandData
	node.mutexForDataRequest.Unlock()

	signatureBuffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(signatureBuffer)
	err = encoder.Encode(signatureForETH)
	if err != nil {
		logger.Error("Failed To encode Signature For Ethereum", "Error", err)
		return fmt.Errorf("something went wrong while encoding signature for ethereum: %v", err)
	}
	res.SignatureForETH = hex.EncodeToString(signatureBuffer.Bytes())
	logger.Info("Finalized RandData Is Approved And Signature For Ethereum Is Shared With Leader Peer", "ResData", res)

	return nil
}

func (node *Node) ReceiveFinalizedRandDataSupportInfo(r *http.Request, req *ReceiveFinalizedRandDataSupportInfoArgs, res *EmptyReply) error {
	logger.Info("Received `ReceiveFinalizedRandDataSupportInfo` Request", "RemoteAddr", r.RemoteAddr)
	leaderPeerIp, _, _ := net.SplitHostPort(strings.TrimPrefix(node.PeersAddr[node.PeerList[node.LeaderPeerIndex]].RequestAddr, "http://"))
	requesterIp, _, _ := net.SplitHostPort(r.RemoteAddr)
	if requesterIp != leaderPeerIp && requesterIp != "127.0.0.1" {
		logger.Error("Received Invalid Request", "RemoteAddr", r.RemoteAddr)
		return fmt.Errorf("only leader node can send request for storing data")
	}
	logger.Info("Received Data From Node", "Data", *req)

	leaderPeerID := node.PeerList[node.LeaderPeerIndex]
	if err := node.verifySignature(leaderPeerID, req.LeaderSignature, req.Nonce, PeerSignatureMessage{
		PlatformName: constants.PlatformName,
		Nonce:        req.Nonce,
		Data:         req.TxHash,
	}); err != nil {
		return err
	}
	contract, err := requestManager.NewRequestManager(common.HexToAddress(constants.RequestMangerContract), node.ethClientRPC)
	if err != nil {
		logger.Error("Failed To Create Request Fulfiller Contract Instance", "Error", err)
		return fmt.Errorf("failed to create fulfiller contract instance: %v", err)
	}
	node.mutexForDataRequest.RLock()
	requestReceiver := node.DataForRequest[req.TxHash].Receiver
	requestNum := node.DataForRequest[req.TxHash].RequestNum
	finalizedRandData := node.DataForRequest[req.TxHash].FinalizedRandData
	node.mutexForDataRequest.RUnlock()
	messageX, messageY, err := contract.GetMessageForRequestFulFillmentSig(nil, requestReceiver, requestNum, finalizedRandData)
	if err != nil {
		logger.Error("Failed To Get Signature Message", "Error", err)
		return fmt.Errorf("something went wrong while getting signature message: %v", err)
	}
	for peerId, data := range req.SupportForFinalizedRandData {
		if node.Id != peerId {
			if ok, _ := node.bn128_bls.VerifySignature(data.SignatureForETH, node.PeersRegisteredId[peerId].G2, hex.EncodeToString(messageX), hex.EncodeToString(messageY)); !ok {
				logger.Error("Failed To Verify Signature Of Request Supporter Peer", "PeerId", peerId, "Signature", data.SignatureForETH)
			}
		}
	}
	node.mutexForDataRequest.Lock()
	node.DataForRequest[req.TxHash].SupportForFinalizedRandData = req.SupportForFinalizedRandData
	node.mutexForDataRequest.Unlock()
	logger.Info("Finalized RandData Support Info Stored")

	return nil
}

func (node *Node) ReceiveFulFillRequestTxHash(r *http.Request, req *ReceiveFulFillRequestTxHashArgs, res *EmptyReply) error {
	logger.Info("Received `ReceiveFulFillRequestTxHash` Request", "RemoteAddr", r.RemoteAddr)
	leaderPeerIp, _, _ := net.SplitHostPort(strings.TrimPrefix(node.PeersAddr[node.PeerList[node.LeaderPeerIndex]].RequestAddr, "http://"))
	requesterIp, _, _ := net.SplitHostPort(r.RemoteAddr)
	if requesterIp != leaderPeerIp && requesterIp != "127.0.0.1" {
		logger.Error("Received Invalid Request", "RemoteAddr", r.RemoteAddr)
		return fmt.Errorf("only leader node can send request for storing data")
	}
	logger.Info("Received Data From Node", "Data", *req)

	leaderPeerID := node.PeerList[node.LeaderPeerIndex]
	if err := node.verifySignature(leaderPeerID, req.LeaderSignature, req.Nonce, PeerSignatureMessage{
		PlatformName: constants.PlatformName,
		Nonce:        req.Nonce,
		Data: RequestAndFulFullTxHash{
			RequestTxHash:      req.RequestTxHash,
			FulFillRequestHash: req.FulFillRequestHash,
		},
	}); err != nil {
		return err
	}
	node.mutexForDataRequest.Lock()
	node.DataForRequest[req.RequestTxHash].FulFillTXHash = req.FulFillRequestHash
	node.mutexForDataRequest.Unlock()
	logger.Info("RandData Request Is FullFilled", "RequestTXHash", req.RequestTxHash, "FulFillTxHash", req.FulFillRequestHash)

	return nil
}

func (node *Node) AgreeToRemovePeerFromNodeManager(r *http.Request, req *AgreeToRemovePeerFromNodeManagerArgs, res *AgreeToRemovePeerFromNodeManagerReply) error {
	logger.Info("Received `AgreeToRemovePeerFromNodeManager` Request", "RemoteAddr", r.RemoteAddr)
	leaderPeerIp, _, _ := net.SplitHostPort(strings.TrimPrefix(node.PeersAddr[node.PeerList[node.LeaderPeerIndex]].RequestAddr, "http://"))
	requesterIp, _, _ := net.SplitHostPort(r.RemoteAddr)
	if requesterIp != leaderPeerIp && requesterIp != "127.0.0.1" {
		logger.Error("Received Invalid Request", "RemoteAddr", r.RemoteAddr)
		return fmt.Errorf("only leader node can send request for storing data")
	}
	logger.Info("Received Data From Node", "Data", *req)

	leaderPeerID := node.PeerList[node.LeaderPeerIndex]
	if err := node.verifySignature(leaderPeerID, req.LeaderSignature, req.Nonce, PeerSignatureMessage{
		PlatformName: constants.PlatformName,
		Nonce:        req.Nonce,
		Data:         req.NodeToRemoveHashId,
	}); err != nil {
		return err
	}
	contract, err := nodeManager.NewNodeManager(common.HexToAddress(constants.NodeManagerContract), node.ethClientRPC)
	if err != nil {
		logger.Error("Failed To Create Request Manager Contract Instance", "Error", err)
		return fmt.Errorf("failed to create nodeManager contract instance: %v", err)
	}
	messageX, messageY, err := contract.GetMessageForNodeRemovalSig(nil, req.NodeToRemoveHashId)
	if err != nil {
		logger.Error("Failed To Get Signature Message", "Error", err)
		return fmt.Errorf("failed to get signature message: %v", err)
	}
	signatureToRemovePeer, err := node.bn128_bls.GenerateSignature(node.nodeRegisteredIdKeyPair, hex.EncodeToString(messageX), hex.EncodeToString(messageY))
	if err != nil {
		logger.Error("Failed To Generate Signature To Remove Peer", "Error", err)
		return fmt.Errorf("failed to generate signature to remove peer: %v", err)
	}

	signatureBuffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(signatureBuffer)
	err = encoder.Encode(signatureToRemovePeer)
	if err != nil {
		logger.Error("Failed To encode Signature For Ethereum", "Error", err)
		return fmt.Errorf("something went wrong while encoding signature for ethereum: %v", err)
	}
	res.SignatureToRemovePeer = hex.EncodeToString(signatureBuffer.Bytes())
	node.PeersParticipationRequest.Delete(req.NodeToRemoveHashId)
	logger.Info("Peer Removed Form PeerParticipationRequest List Signature To Remove Peer From NodeManager Is Generated", "NodeToRemoveHashId", req.NodeToRemoveHashId, "SignatureToRemoveNode", res.SignatureToRemovePeer)
	return nil
}
