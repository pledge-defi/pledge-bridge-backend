package abifile

import (
	"os"
	"path"
	"runtime"
)

func GetPledgeBridgeBscAbi() (string, error) {
	currentAbPath := GetCurrentAbPathByCaller()
	by, err := os.ReadFile(currentAbPath + "/pledge_bridge_bsc.abi")
	if err != nil {
		return "", err
	}
	return string(by), nil
}

func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
