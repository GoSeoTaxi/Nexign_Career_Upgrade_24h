package spel

import (
	"Nexign_Career_Upgrade_24h/internal/config"
	"Nexign_Career_Upgrade_24h/internal/etc"
	logger "Nexign_Career_Upgrade_24h/internal/loger"
)

func LimitReq(arT []string) []etc.AnsJ {

	//	fmt.Println(`++++!!!`)
	//	fmt.Println(arT)
	//	fmt.Println(`++++!!!`)

	var StrOutRet []etc.AnsJ
	var err error

	var ReqSpel string
	for i, t1 := range arT {

		if len(t1) > config.SizeReq {
			logger.Infof("err long line")
		}

		if len(ReqSpel+" "+t1) < config.SizeReq {
			ReqSpel = ReqSpel + " " + t1
		} else {
			StrOutRet, err = spellStr(ReqSpel, StrOutRet)
			if err != nil {
				logger.Infof("err RequestSpel - %v", err)
			}
			ReqSpel = t1
		}
		if i+1 == len(arT) {
			StrOutRet, err = spellStr(ReqSpel, StrOutRet)
			if err != nil {
				logger.Infof("err RequestSpel - %v", err)
			}
		}
		i++
	}

	return StrOutRet
}
