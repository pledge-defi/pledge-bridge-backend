package validate

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-bridge-backend/api/common/statecode"
	"pledge-bridge-backend/api/models/request"
)

type TxHistory struct{}

func NewTxHistory() *TxHistory {
	return &TxHistory{}
}

func (s *TxHistory) TxHistory(c *gin.Context, req *request.TxHistory) int {

	err := c.ShouldBindJSON(req)
	fmt.Println(err, 88888)
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "Address" && e.Tag() == "required" {
				return statecode.AddressEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	return statecode.CommonSuccess
}
