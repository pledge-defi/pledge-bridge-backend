package controllers

import (
	"github.com/gin-gonic/gin"
	"pledge-bridge-backend/api/common/statecode"
	"pledge-bridge-backend/api/models/request"
	"pledge-bridge-backend/api/models/response"
	"pledge-bridge-backend/api/services"
	"pledge-bridge-backend/api/validate"
)

type TransitionController struct {
}

func (c *TransitionController) TxHistory(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.TxHistory{}
	result := response.TxHistory{}

	errCode := validate.NewTxHistory().TxHistory(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	errCode, count, txHistory := services.NewTxHistory().TxHistory(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	result.Rows = txHistory
	result.Count = count
	res.Response(ctx, statecode.CommonSuccess, result)
	return
}

func (c *TransitionController) LockedCountdown(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	result := response.LockedCountdown{}

	errCode := services.NewLockedCountdown().NextTimeSeconds(&result)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, result)
	return
}

func (c *TransitionController) AddTx(ctx *gin.Context) {
	req := request.AddTx{}
	res := response.Gin{Res: ctx}

	errCode := validate.NewAddTx().TxInfo(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	errCode = services.NewAddTxService().AddRecord(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, nil)
	return
}
