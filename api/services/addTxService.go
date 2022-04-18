package services

import (
	"pledge-bridge-backend/api/common/statecode"
	"pledge-bridge-backend/api/models"
	"pledge-bridge-backend/api/models/request"
	"pledge-bridge-backend/log"
)

type AddTxService struct{}

func NewAddTxService() *AddTxService {
	return &AddTxService{}
}

func (c *AddTxService) AddRecord(req *request.AddTx) int {

	err := models.NewTxHistory().AddRecord(req)
	if err != nil {
		log.Logger.Error(err.Error())
		return statecode.CommonErrServerErr
	}
	return statecode.CommonSuccess
}
