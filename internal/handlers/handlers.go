package handlers

import (
	logger "Nexign_Career_Upgrade_24h/internal/loger"
	"Nexign_Career_Upgrade_24h/internal/spel"
	"encoding/json"
	"io"
	"net/http"
)

func GetAPIJSON(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	type TJSON struct {
		Texts []string `json:"texts"`
	}

	var ArJSON TJSON
	err = json.Unmarshal([]byte(b), &ArJSON)
	if err != nil {
		logger.Infof("UnMarshall in - %v", err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(spel.Makeout(spel.LimitReq(ArJSON.Texts)))

}
