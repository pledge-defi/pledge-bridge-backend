package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"math/big"
	"pledge-bridge-backend/api/common/statecode"
	"pledge-bridge-backend/api/models/request"
)

type AddTx struct{}

func NewAddTx() *AddTx {
	return &AddTx{}
}

func (s *AddTx) TxInfo(c *gin.Context, req *request.AddTx) int {

	err := c.ShouldBind(req)
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "Address" && e.Tag() == "required" {
				return statecode.AddressEmpty
			}
			if e.Field() == "TxType" && e.Tag() == "required" {
				return statecode.TxTypeEmpty
			}
			if e.Field() == "Fee" && e.Tag() == "required" {
				return statecode.FeeEmpty
			}
			if e.Field() == "Asset" && e.Tag() == "required" {
				return statecode.AssetEmpty
			}
			if e.Field() == "TxHash" && e.Tag() == "required" {
				return statecode.TxHashEmpty
			}
			if e.Field() == "Amount" && e.Tag() == "required" {
				return statecode.AmountEmpty
			}
			if e.Field() == "SrcChain" && e.Tag() == "required" {
				return statecode.SrcChainEmpty
			}
			if e.Field() == "DestChain" && e.Tag() == "required" {
				return statecode.DestChainEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	bigIntAmount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		return statecode.AmountEmpty
	}
	bigIntZero, _ := new(big.Int).SetString("0", 10)
	if bigIntAmount.Cmp(bigIntZero) <= 0 {
		return statecode.AmountEmpty
	}

	bigIntAmount, ok = new(big.Int).SetString(req.Fee, 10)
	if !ok {
		return statecode.FeeEmpty
	}
	if bigIntAmount.Cmp(bigIntZero) <= 0 {
		return statecode.FeeEmpty
	}

	return statecode.CommonSuccess
}
