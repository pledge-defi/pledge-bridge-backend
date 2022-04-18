package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-bridge-backend/api/common/statecode"
	"pledge-bridge-backend/api/models/request"
)

type UserAssets struct{}

func NewUserAssets() *UserAssets {
	return &UserAssets{}
}

func (s *UserAssets) UserAssets(c *gin.Context, req *request.UserAssets) int {

	err := c.ShouldBind(req)
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "Token" && e.Tag() == "required" {
				return statecode.TokenEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	return statecode.CommonSuccess
}
