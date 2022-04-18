package response

import "pledge-bridge-backend/api/models"

type TxHistory struct {
	Count int64              `json:"count"`
	Rows  []models.TxHistory `json:"rows"`
}
