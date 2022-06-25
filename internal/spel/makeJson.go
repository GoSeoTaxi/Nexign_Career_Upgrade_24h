package spel

import (
	"Nexign_Career_Upgrade_24h/internal/etc"
	logger "Nexign_Career_Upgrade_24h/internal/loger"
	"encoding/json"
)

func Makeout(StrOut []etc.AnsJ) []byte {

	b, err := json.Marshal(StrOut)
	if err != nil {
		logger.Infof("err MarshalOut - %v", err)
	}

	return b
}
