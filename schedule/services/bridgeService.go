package services

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"pledge-bridge-backend/contract/bindings/bridgeBSC"
	"pledge-bridge-backend/log"
	"pledge-bridge-backend/schedule/models"
	"time"
)

type BridgeService struct {
}

func NewBridgeService() *BridgeService {
	return &BridgeService{}
}

// BridgeUpKeep  release PLGR
func (s *BridgeService) BridgeUpKeep(netUrl, bridgeBSCToken, privateKey string, chainId int64) {

	ethereumClient, err := ethclient.Dial(netUrl)
	if nil != err {
		log.Logger.Error(err.Error())
		return
	}
	defer ethereumClient.Close()

	bridgeBSCContract, err := bridgeBSC.NewTokengo(common.HexToAddress(bridgeBSCToken), ethereumClient)
	if nil != err {
		log.Logger.Error(err.Error())
		return
	}

	privateKeyEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	// No keystone file, so discard ; The following method returns the next method, so use the next method directly
	//auth, err := bind.NewTransactorWithChainID(file, "pledge bridge", big.NewInt(97))
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(chainId))
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	transactOpts := bind.TransactOpts{
		From:      auth.From,
		Nonce:     nil,
		Signer:    auth.Signer, // Method to use for signing the transaction (mandatory)
		Value:     big.NewInt(0),
		GasPrice:  nil,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  0,
		Context:   ctx,
		NoSend:    false, // Do all transact steps but do not send the transaction
	}

	// Underprice and other situations may occur during execution. Retry can solve this problem
	retry := 5
	tx := &types.Transaction{}
	for {
		tx, err = bridgeBSCContract.ExecuteUpkeep(&transactOpts)
		log.Logger.Sugar().Info("tx info ", tx)
		if err != nil {
			log.Logger.Error(err.Error())
		}
		if retry <= 0 || err == nil {
			break
		}
		retry--
		time.Sleep(time.Minute)
	}

	if err == nil {
		log.Logger.Info("release PLGR success")
		// update db
		where := fmt.Sprintf("bridge_hash='' and tx_type='0' and asset='PLGR'")
		err = models.NewTxHistory().UpdateBridgeHash(where, tx.Hash().String(), tx.Cost().String())
		if err != nil {
			log.Logger.Error(err.Error())
			return
		}
		log.Logger.Info("release and update db success")
	}

}

// BridgeUpKeepFailed execute failed
func (s *BridgeService) BridgeUpKeepFailed() {

	fmt.Println("多次尝试后跨链仍然失败")
}

// KeystoreToPrivateKey import privateKey from keystone file
func KeystoreToPrivateKey(privateKeyFile, password string) (string, string, error) {
	keyjson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Logger.Error(err.Error())
		return "", "", err
	}
	unlockedKey, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		log.Logger.Error(err.Error())
		return "", "", err
	}
	privKey := hex.EncodeToString(unlockedKey.PrivateKey.D.Bytes())
	addr := crypto.PubkeyToAddress(unlockedKey.PrivateKey.PublicKey)
	return privKey, addr.String(), nil
}
