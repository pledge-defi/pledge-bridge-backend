package controllers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	apicommon "pledge-bridge-backend/api/common"
	"pledge-bridge-backend/api/common/statecode"
	"pledge-bridge-backend/api/models/request"
	"pledge-bridge-backend/api/models/response"
	"pledge-bridge-backend/api/services"
	"pledge-bridge-backend/api/validate"
	"pledge-bridge-backend/config"
	"pledge-bridge-backend/contract/bindings/bridgeBSC"
	"pledge-bridge-backend/contract/bindings/bridgeETH"
	"pledge-bridge-backend/log"
)

type UserAssetsController struct {
	LockedPlgr       string
	PlgrCanWithdraw  string
	MplgrCanWithdraw string
}

func (c *UserAssetsController) UserAssets(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.UserAssets{}
	result := response.UserAssets{}
	eg := errgroup.Group{}

	errCode := validate.NewUserAssets().UserAssets(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	eg.Go(func() error {
		return c.UserBscAssets(req.Token)
	})

	eg.Go(func() error {
		return c.UserEthAssets(req.Token)
	})

	//Asynchronous update database
	apicommon.UserAssetChain <- req.Token

	if err := eg.Wait(); err != nil {
		log.Logger.Sugar().Error("web3 called failed, get data from db ", err)
		errCode := services.NewUserAssetsService().UserAssets(&req, &result)
		if errCode != statecode.CommonSuccess {
			res.Response(ctx, errCode, nil)
			return
		}
		res.Response(ctx, statecode.CommonSuccess, result)
		return
	}

	result.Token = req.Token
	result.PlgrCanWithdraw = c.PlgrCanWithdraw
	result.MplgrCanWithdraw = c.MplgrCanWithdraw
	result.LockedPlgr = c.LockedPlgr

	res.Response(ctx, statecode.CommonSuccess, result)
	return
}

func (c *UserAssetsController) UserBscAssets(token string) error {

	ethereumClient, err := ethclient.Dial(config.Config.BscNet.NetUrl)
	if nil != err {
		log.Logger.Error(err.Error())
		return err
	}
	defer ethereumClient.Close()

	bridgeBSCContract, err := bridgeBSC.NewTokengo(common.HexToAddress(config.Config.BscNet.BridgeToken), ethereumClient)
	if nil != err {
		log.Logger.Error(err.Error())
		return err
	}
	plgrCanWithdraw, err := bridgeBSCContract.PlgrAmounts(nil, common.HexToAddress(token))
	if nil != err {
		log.Logger.Error(err.Error())
		return err
	}

	lockedPlgr, err := bridgeBSCContract.LockedPlgrTx(nil, common.HexToAddress(token))
	if nil != err {
		log.Logger.Error(err.Error())
		return err
	}

	c.LockedPlgr = lockedPlgr.Amount.String()
	c.PlgrCanWithdraw = plgrCanWithdraw.String()

	return nil

}

func (c *UserAssetsController) UserEthAssets(token string) error {

	ethereumClient, err := ethclient.Dial(config.Config.EthNet.NetUrl)
	if nil != err {
		log.Logger.Error(err.Error())
		return err
	}
	defer ethereumClient.Close()

	bridgeETHContract, err := bridgeETH.NewTokengo(common.HexToAddress(config.Config.EthNet.BridgeToken), ethereumClient)
	if nil != err {
		log.Logger.Error(err.Error())
		return err
	}
	mplgrCanWithdraw, err := bridgeETHContract.MplgrAmounts(nil, common.HexToAddress(token))
	if nil != err {
		log.Logger.Error(err.Error())
		return err
	}

	c.MplgrCanWithdraw = mplgrCanWithdraw.String()

	return nil

}
