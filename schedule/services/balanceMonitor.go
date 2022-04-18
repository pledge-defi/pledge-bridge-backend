package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"math/big"
	"pledge-bridge-backend/config"
	"pledge-bridge-backend/contract/bindings/mplgr"
	"pledge-bridge-backend/contract/bindings/plgr"
	"pledge-bridge-backend/log"
	"pledge-bridge-backend/utils"
)

type BalanceMonitor struct {
}

func NewBalanceMonitor() *BalanceMonitor {
	return &BalanceMonitor{}
}

// Monitor Sending email when balance is insufficient
func (s *BalanceMonitor) Monitor() {

	// check bridge-bsc-contract PLGR
	bridgeBSCBalance, err := s.GetPlgrBalance()
	thresholdBridgeBSC, ok := new(big.Int).SetString(config.Config.Threshold.BridgeBsc, 10)
	if ok && (err == nil) && (bridgeBSCBalance.Cmp(thresholdBridgeBSC) <= 0) {
		emailBody, err := s.EmailBody(config.Config.Contract.PlgrToken, "PLGR", bridgeBSCBalance.String(), thresholdBridgeBSC.String())
		if err != nil {
			log.Logger.Error(err.Error())
		} else {
			err = utils.SendEmail(emailBody, 2)
			if err != nil {
				log.Logger.Error(err.Error())
			}
		}
	}

	// check  bridge-eth-contract MPLGR
	bridgeETHBalance, err := s.GetMplgrBalance()
	thresholdBridgeETH, ok := new(big.Int).SetString(config.Config.Threshold.BridgeEth, 10)
	if ok && (err == nil) && (bridgeETHBalance.Cmp(thresholdBridgeETH) <= 0) {
		emailBody, err := s.EmailBody(config.Config.Contract.MplgrToken, "MPLGR", bridgeETHBalance.String(), thresholdBridgeETH.String())
		if err != nil {
			log.Logger.Error(err.Error())
		} else {
			err = utils.SendEmail(emailBody, 2)
			if err != nil {
				log.Logger.Error(err.Error())
			}
		}
	}

	//check bridge account bsc
	tokenOneBalance, err := s.GetBalance(config.Config.BscNet.NetUrl, config.Config.Contract.BridgeTokenOne)
	thresholdBridgeTokenOne, ok := new(big.Int).SetString(config.Config.Threshold.BridgeTokenOneBnb, 10)
	if ok && (err == nil) && (tokenOneBalance.Cmp(thresholdBridgeTokenOne) <= 0) {
		emailBody, err := s.EmailBody(config.Config.Contract.BridgeTokenOne, "BNB", tokenOneBalance.String(), thresholdBridgeTokenOne.String())
		if err != nil {
			log.Logger.Error(err.Error())
		} else {
			err = utils.SendEmail(emailBody, 2)
			if err != nil {
				log.Logger.Error(err.Error())
			}
		}
	}
	tokenTwoBalance, err := s.GetBalance(config.Config.BscNet.NetUrl, config.Config.Contract.BridgeTokenTwo)
	thresholdBridgeTokenTwo, ok := new(big.Int).SetString(config.Config.Threshold.BridgeTokenTwoBnb, 10)
	if ok && (err == nil) && (tokenTwoBalance.Cmp(thresholdBridgeTokenTwo) <= 0) {
		emailBody, err := s.EmailBody(config.Config.Contract.BridgeTokenTwo, "BNB", tokenTwoBalance.String(), thresholdBridgeTokenTwo.String())
		if err != nil {
			log.Logger.Error(err.Error())
		} else {
			err = utils.SendEmail(emailBody, 2)
			if err != nil {
				log.Logger.Error(err.Error())
			}
		}
	}
	tokenThreeBalance, err := s.GetBalance(config.Config.BscNet.NetUrl, config.Config.Contract.BridgeTokenThree)
	thresholdBridgeTokenThree, ok := new(big.Int).SetString(config.Config.Threshold.BridgeTokenThreeBnb, 10)
	if ok && (err == nil) && (tokenThreeBalance.Cmp(thresholdBridgeTokenThree) <= 0) {
		emailBody, err := s.EmailBody(config.Config.Contract.BridgeTokenThree, "BNB", tokenThreeBalance.String(), thresholdBridgeTokenThree.String())
		if err != nil {
			log.Logger.Error(err.Error())
		} else {
			err = utils.SendEmail(emailBody, 2)
			if err != nil {
				log.Logger.Error(err.Error())
			}
		}
	}

	//check bridge account eth
	tokenOneBalance, err = s.GetBalance(config.Config.EthNet.NetUrl, config.Config.Contract.BridgeTokenOne)
	thresholdBridgeTokenOne, ok = new(big.Int).SetString(config.Config.Threshold.BridgeTokenOneEth, 10)
	if ok && (err == nil) && (tokenOneBalance.Cmp(thresholdBridgeTokenOne) <= 0) {
		emailBody, err := s.EmailBody(config.Config.Contract.BridgeTokenOne, "ETH", tokenOneBalance.String(), thresholdBridgeTokenOne.String())
		if err != nil {
			log.Logger.Error(err.Error())
		} else {
			err = utils.SendEmail(emailBody, 2)
			if err != nil {
				log.Logger.Error(err.Error())
			}
		}
	}
	tokenTwoBalance, err = s.GetBalance(config.Config.EthNet.NetUrl, config.Config.Contract.BridgeTokenTwo)
	thresholdBridgeTokenTwo, ok = new(big.Int).SetString(config.Config.Threshold.BridgeTokenTwoEth, 10)
	if ok && (err == nil) && (tokenTwoBalance.Cmp(thresholdBridgeTokenTwo) <= 0) {
		emailBody, err := s.EmailBody(config.Config.Contract.BridgeTokenTwo, "ETH", tokenTwoBalance.String(), thresholdBridgeTokenTwo.String())
		if err != nil {
			log.Logger.Error(err.Error())
		} else {
			err = utils.SendEmail(emailBody, 2)
			if err != nil {
				log.Logger.Error(err.Error())
			}
		}
	}
	tokenThreeBalance, err = s.GetBalance(config.Config.EthNet.NetUrl, config.Config.Contract.BridgeTokenThree)
	thresholdBridgeTokenThree, ok = new(big.Int).SetString(config.Config.Threshold.BridgeTokenThreeEth, 10)
	if ok && (err == nil) && (tokenThreeBalance.Cmp(thresholdBridgeTokenThree) <= 0) {
		emailBody, err := s.EmailBody(config.Config.Contract.BridgeTokenThree, "ETH", tokenThreeBalance.String(), thresholdBridgeTokenThree.String())
		if err != nil {
			log.Logger.Error(err.Error())
		} else {
			err = utils.SendEmail(emailBody, 2)
			if err != nil {
				log.Logger.Error(err.Error())
			}
		}
	}
}

// GetBalance get balance of ERC20 token
func (s *BalanceMonitor) GetBalance(netUrl, token string) (*big.Int, error) {

	ethereumClient, err := ethclient.Dial(netUrl)
	if err != nil {
		log.Logger.Error(err.Error())
		return big.NewInt(0), err
	}
	defer ethereumClient.Close()

	balance, err := ethereumClient.BalanceAt(context.Background(), common.HexToAddress(token), nil)
	if err != nil {
		log.Logger.Error(err.Error())
		return big.NewInt(0), err
	}

	return balance, err
}

// GetPlgrBalance get balance of PLGR
func (s *BalanceMonitor) GetPlgrBalance() (*big.Int, error) {
	ethereumClient, err := ethclient.Dial(config.Config.BscNet.NetUrl)
	if err != nil {
		log.Logger.Error(err.Error())
		return big.NewInt(0), err
	}
	defer ethereumClient.Close()

	contract, err := plgr.NewTokengo(common.HexToAddress(config.Config.Contract.PlgrToken), ethereumClient)
	if err != nil {
		log.Logger.Error(err.Error())
		return big.NewInt(0), err
	}

	balance, err := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(config.Config.BscNet.BridgeToken))
	if err != nil {
		log.Logger.Error(err.Error())
		return big.NewInt(0), err
	}
	return balance, err
}

// GetMplgrBalance get balance of MPLGR
func (s *BalanceMonitor) GetMplgrBalance() (*big.Int, error) {
	ethereumClient, err := ethclient.Dial(config.Config.EthNet.NetUrl)
	if err != nil {
		log.Logger.Error(err.Error())
		return big.NewInt(0), err
	}
	defer ethereumClient.Close()

	contract, err := mplgr.NewTokengo(common.HexToAddress(config.Config.Contract.MplgrToken), ethereumClient)
	if err != nil {
		log.Logger.Error(err.Error())
		return big.NewInt(0), err
	}

	balance, err := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(config.Config.EthNet.BridgeToken))
	if err != nil {
		log.Logger.Error(err.Error())
		return big.NewInt(0), err
	}
	return balance, err
}

// EmailBody email body
func (s *BalanceMonitor) EmailBody(token, currency, balance, threshold string) ([]byte, error) {
	e18, err := decimal.NewFromString("1000000000000000000")
	if err != nil {
		return nil, err
	}

	balanceDeci, err := decimal.NewFromString(balance)
	if err != nil {
		return nil, err
	}
	balanceStr := balanceDeci.Div(e18).String()

	thresholdDeci, err := decimal.NewFromString(threshold)
	if err != nil {
		return nil, err
	}

	thresholdStr := thresholdDeci.Div(e18).String()
	log.Logger.Sugar().Info("balance not enough ", token, " ", currency, " ", balanceStr, " ", thresholdStr)
	body := fmt.Sprintf(`<p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The balance of <strong><span style="color: rgb(255, 0, 0);"> %s </span></strong> is <strong>%s %s </strong>. Please recharge it in time. The current minimum balance limit is %s %s 
</p>`, token, balanceStr, currency, thresholdStr, currency)
	return []byte(body), nil
}
