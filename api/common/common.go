package common

import (
	"pledge-bridge-backend/config"
	"time"
)

var UserAssetChain = make(chan string, 100)

var ReleaseTime int64

func SetReleaseTime(sec int64) {
	if config.Config.Env.IsDev {
		ReleaseTime = time.Now().Add(time.Second * time.Duration(sec)).Unix()
	}
}
