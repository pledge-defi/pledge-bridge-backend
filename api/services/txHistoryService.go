package services

import (
	"fmt"
	"pledge-bridge-backend/api/common/statecode"
	"pledge-bridge-backend/api/models"
	"pledge-bridge-backend/api/models/request"
	"pledge-bridge-backend/log"
)

type TxHistoryService struct{}

func NewTxHistory() *TxHistoryService {
	return &TxHistoryService{}
}

func (c *TxHistoryService) TxHistory(req *request.TxHistory) (int, int64, []models.TxHistory) {

	whereCondition := fmt.Sprintf(`address = '%v' and tx_type= %s `, req.Address, req.TxType)

	err, total, data := models.NewTxHistory().Pagination(req, whereCondition)
	if err != nil {
		log.Logger.Error(err.Error())
		return statecode.CommonErrServerErr, 0, nil
	}
	return 0, total, data
}
