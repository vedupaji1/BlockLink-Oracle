package temp11

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"sync"
	bn128_blsPKG "temp11/bn128_bls"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	logger "github.com/inconshreveable/log15"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"github.com/syndtr/goleveldb/leveldb"
)

func StoreStringMap(db *leveldb.DB, key string, value map[string]string) {
	var stringSliceBytes bytes.Buffer
	gobEncoder := gob.NewEncoder(&stringSliceBytes)
	err := gobEncoder.Encode(value)
	if err != nil {
		log.Panic("Failed To Encode String Map: ", err)
	}
	err = db.Put([]byte(key), stringSliceBytes.Bytes(), nil)
	if err != nil {
		log.Panic("Failed To Set String Map: ", err)
	}
}

func GetMapString(db *leveldb.DB, key string, res *map[string]string) {
	stringSliceBytes, err := db.Get([]byte(key), nil)
	if err != nil {
		log.Panic("Failed To Get String Map: ", err)
	}
	gobDecoder := gob.NewDecoder(bytes.NewBuffer(stringSliceBytes))
	err = gobDecoder.Decode(res)
	if err != nil {
		log.Panic("Failed To Decode String Map: ", err)
	}
}

func CloneStringMap(mapToBeCopied map[string]PeerReqAndListenAddr) map[string]PeerReqAndListenAddr {
	res := make(map[string]PeerReqAndListenAddr)
	for key, value := range mapToBeCopied {
		res[key] = value
	}
	return res
}

func HashRegisteredId(registeredId [4]*big.Int) (common.Hash, error) {
	uintArrType, err := abi.NewType("uint256[4]", "", nil)
	if err != nil {
		logger.Error("Failed To Generate uint256 array type", "Error", err)
		return common.Hash{}, fmt.Errorf("failed to generate uint256 array type: %v", err)
	}
	arguments := abi.Arguments{
		{
			Type: uintArrType,
		},
	}
	pubKeyBytes, err := arguments.Pack(registeredId)
	if err != nil {
		logger.Error("Failed To Convert Array To Bytes", "Error", err)
		return common.Hash{}, fmt.Errorf("failed to convert array to bytes: %v", err)
	}
	return crypto.Keccak256Hash(pubKeyBytes), nil
}

func GeneratePeerIndex(bn128_bls *bn128_blsPKG.BLS, peerList []string, peersRegisteredId map[string]PeersRegisteredId) map[common.Hash]int {
	res := make(map[common.Hash]int)
	for index, data := range peerList {
		peerRegisteredIdHash, err := HashRegisteredId(bn128_bls.ParsePubKey((peersRegisteredId[data].G2)))
		if err != nil {
			logger.Error("", "Error", err)
			log.Panic(err)
		}
		res[peerRegisteredIdHash] = index
	}
	return res
}

func ParseBootstrappingPeerConfigData(bn128_bls *bn128_blsPKG.BLS, mapToBeConverted map[string]interface{}) (map[string]PeerReqAndListenAddr, map[string]PeersRegisteredId) {
	peersAddr := make(map[string]PeerReqAndListenAddr)
	peersRegisteredId := make(map[string]PeersRegisteredId)
	for key, value := range mapToBeConverted {
		peersData := struct {
			ListenAddr     string
			RequestAddr    string
			RegisteredIdG1 string
			RegisteredIdG2 string
		}{}
		err := mapstructure.Decode(value, &peersData)
		if err != nil {
			log.Panic("Failed To Decode Config Data: ", err)
		}
		peersAddr[key] = PeerReqAndListenAddr{
			ListenAddr:  peersData.ListenAddr,
			RequestAddr: peersData.RequestAddr,
		}
		g1InBytes, err := hex.DecodeString(peersData.RegisteredIdG1)
		if err != nil {
			log.Panic("Failed To Decode Hex String Data: ", err)
		}
		g2InBytes, err := hex.DecodeString(peersData.RegisteredIdG2)
		if err != nil {
			log.Panic("Failed To Decode Hex String Data: ", err)
		}

		decodedG1 := [3]*big.Int{}
		decoderG1 := gob.NewDecoder(bytes.NewBuffer(g1InBytes))
		decoderG1.Decode(&decodedG1)

		decodedG2 := [3][2]*big.Int{}
		decoderG2 := gob.NewDecoder(bytes.NewBuffer(g2InBytes))
		decoderG2.Decode(&decodedG2)
		peersRegisteredId[key] = PeersRegisteredId{
			G1: decodedG1,
			G2: decodedG2,
		}
	}
	return peersAddr, peersRegisteredId
}

func GenBigIntRand() (*big.Int, error) {
	max := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(130), nil)
	randNum, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random number: %v", err)
	}
	return randNum, nil
}

func SyncMapToStringIntMap(syncMap *sync.Map) map[string]int {
	res := make(map[string]int)
	syncMap.Range(func(key, value any) bool {
		res[cast.ToString(key)] = cast.ToInt(value)
		return true
	})
	return res
}

func StringIntMapToSyncMap(stringIntMap map[string]int) *sync.Map {
	res := new(sync.Map)
	for key, value := range stringIntMap {
		res.Store(key, value)
	}
	return res
}

func CopyStringArray(arr []string) []string {
	return append([]string{}, arr...)
}
