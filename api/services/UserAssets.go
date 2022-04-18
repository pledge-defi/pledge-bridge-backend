package services

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	apicommon "pledge-bridge-backend/api/common"
	"pledge-bridge-backend/api/common/statecode"
	"pledge-bridge-backend/api/models"
	"pledge-bridge-backend/api/models/request"
	"pledge-bridge-backend/api/models/response"
	"pledge-bridge-backend/config"
	"pledge-bridge-backend/contract/bindings/bridgeBSC"
	"pledge-bridge-backend/contract/bindings/bridgeETH"
	"pledge-bridge-backend/log"
	"pledge-bridge-backend/utils"
	"sync"
	"time"
)

type UserAssetsService struct{}

func NewUserAssetsService() *UserAssetsService {
	return &UserAssetsService{}
}

func (c *UserAssetsService) UserAssets(req *request.UserAssets, res *response.UserAssets) int {

	err, result := models.NewUserAssets().GetInfoByToken(req)
	if err != nil {
		log.Logger.Error(err.Error())
		return statecode.CommonErrServerErr
	}

	res.Token = result.Token
	res.PlgrCanWithdraw = result.PlgrCanWithdraw
	res.MplgrCanWithdraw = result.MplgrCanWithdraw
	res.LockedPlgr = result.LockedPlgr

	return statecode.CommonSuccess
}

func (c *UserAssetsService) UserBscAssets(ctx context.Context, token string) {

	var done = make(chan bool)
	go func() {
		ethereumClient, err := ethclient.Dial(config.Config.BscNet.NetUrl)
		if nil != err {
			log.Logger.Error(err.Error())
			done <- true
			return
		}
		defer ethereumClient.Close()

		bridgeBSCContract, err := bridgeBSC.NewTokengo(common.HexToAddress(config.Config.BscNet.BridgeToken), ethereumClient)
		if nil != err {
			log.Logger.Error(err.Error())
			done <- true
			return
		}
		plgrCanWithdraw, err := bridgeBSCContract.PlgrAmounts(nil, common.HexToAddress(token))
		if nil != err {
			log.Logger.Error(err.Error())
			done <- true
			return
		}

		lockedPlgr, err := bridgeBSCContract.LockedPlgrTx(nil, common.HexToAddress(token))
		if nil != err {
			log.Logger.Error(err.Error())
			done <- true
			return
		}

		err = models.NewUserAssets().UpdateUserAssets(token, map[string]interface{}{
			"locked_plgr":       lockedPlgr.Amount.String(),
			"plgr_can_withdraw": plgrCanWithdraw.String(),
			"updated_at":        utils.NowDataTime(),
		})
		if nil != err {
			log.Logger.Error(err.Error())
			done <- true
			return
		}
		done <- true
	}()

	select {
	case <-ctx.Done():
		log.Logger.Error("timeout")
		return
	case <-done:
		log.Logger.Info("done")
		return
	}
}

func (c *UserAssetsService) UserEthAssets(ctx context.Context, token string) {

	var done = make(chan bool)
	go func() {
		ethereumClient, err := ethclient.Dial(config.Config.EthNet.NetUrl)
		if nil != err {
			log.Logger.Error(err.Error())
			done <- true
			return
		}
		defer ethereumClient.Close()

		bridgeETHContract, err := bridgeETH.NewTokengo(common.HexToAddress(config.Config.EthNet.BridgeToken), ethereumClient)
		if nil != err {
			log.Logger.Error(err.Error())
			done <- true
			return
		}

		mplgrCanWithdraw, err := bridgeETHContract.MplgrAmounts(nil, common.HexToAddress(token))
		if nil != err {
			log.Logger.Error(err.Error())
			done <- true
			return
		}

		err = models.NewUserAssets().UpdateUserAssets(token, map[string]interface{}{
			"mplgr_can_withdraw": mplgrCanWithdraw.String(),
			"updated_at":         utils.NowDataTime(),
		})
		if nil != err {
			log.Logger.Error(err.Error())
			done <- true
			return
		}
		done <- true
	}()

	select {
	case <-ctx.Done():
		log.Logger.Error("timeout")
		return
	case <-done:
		log.Logger.Info("done")
		return
	}
}

func (c *UserAssetsService) UpdateUserAssets() {

	for t := range apicommon.UserAssetChain {

		go func() {
			wg := sync.WaitGroup{}
			wg.Add(2)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			go c.UserBscAssets(ctx, t)
			go c.UserEthAssets(ctx, t)

			wg.Wait()
		}()

	}
}
