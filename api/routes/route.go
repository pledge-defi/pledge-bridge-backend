package routes

import (
	"github.com/gin-gonic/gin"
	"pledge-bridge-backend/api/controllers"
)

func InitRoute(e *gin.Engine) *gin.Engine {

	// version group
	v2Group := e.Group("/api/v2")

	transitionController := controllers.TransitionController{}
	v2Group.POST("/txsHistory", transitionController.TxHistory)            //transitionHistory
	v2Group.POST("/lockedCountdown", transitionController.LockedCountdown) //Time from next release time
	v2Group.POST("/addTx", transitionController.AddTx)                     // add transition record

	userAssetsController := controllers.UserAssetsController{}
	v2Group.GET("/userAssets", userAssetsController.UserAssets) //Get User Assets

	return e
}
