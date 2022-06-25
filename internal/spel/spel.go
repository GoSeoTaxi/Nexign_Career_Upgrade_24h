package spel

import (
	"Nexign_Career_Upgrade_24h/internal/config"
	"Nexign_Career_Upgrade_24h/internal/etc"
	logger "Nexign_Career_Upgrade_24h/internal/loger"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func spellStr(inp string, StrOutRetI []etc.AnsJ) ([]etc.AnsJ, error) {

	i := strings.Index(inp, " ")
	if i == 0 {
		inp = inp[1:len(inp)]
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", config.URLSPELL, nil)
	if err != nil {
		logger.Infof("err RequestSpel - %v", err)
	}
	q := req.URL.Query()
	q.Add("text", inp)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		logger.Infof("err ServerSpel - %v", err)
		return StrOutRetI, err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Infof("err RequestBodySpel - %v", err)
	}

	var outJ []etc.TJSONIn
	err = json.Unmarshal([]byte(responseBody), &outJ)
	if err != nil {
		logger.Infof("err RequestBodySpelNoJSON - %v", err)
	}

	var AArr etc.AnsJ
	for i := 0; i < len(outJ); i++ {
		a := outJ[i]
		for i1 := 0; i1 < len(a); i1++ {
			b := a[i1]
			AArr.Word = b.Word
			AArr.S = b.S
			StrOutRetI = append(StrOutRetI, AArr)
		}
	}

	return StrOutRetI, err

}
