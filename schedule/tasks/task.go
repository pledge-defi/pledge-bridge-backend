package tasks

import (
	"github.com/jasonlvhit/gocron"
	"pledge-bridge-backend/config"
	"pledge-bridge-backend/schedule/common"
	"pledge-bridge-backend/schedule/services"
	"time"
)

// Task bridge task
func Task() {

	// get environment variables
	common.GetEnv()

	scheduler := gocron.NewScheduler()

	scheduler.ChangeLoc(time.UTC)

	_ = scheduler.Every(1).Hour().From(gocron.NextTick()).Do(services.NewBridgeService().BridgeUpKeep, config.Config.BscNet.NetUrl, config.Config.BscNet.BridgeToken, common.BridgeBscAdminPrivateKey, config.Config.BscNet.ChainId) // Begin job immediately upon start

	//_ = scheduler.Every(1).Sunday().At("00:00").Do(services.NewBridgeService().BridgeUpKeep, config.Config.BscNet.NetUrl, config.Config.BscNet.BridgeToken, common.BridgeBscAdminPrivateKey, config.Config.BscNet.ChainId)

	_ = scheduler.Every(20).Minutes().Do(services.NewBalanceMonitor().Monitor) // Sending email when balance is insufficient

	<-scheduler.Start() // Start all the pending jobs

}
