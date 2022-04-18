package common

import (
	"os"
	"pledge-bridge-backend/log"
)

var BridgeBscAdminPrivateKey string

func GetEnv() {

	var ok bool

	BridgeBscAdminPrivateKey, ok = os.LookupEnv("plgr_admin_private_key")
	if !ok {
		log.Logger.Error("environment variable is not set")
		panic("environment variable is not set")
	}

}
