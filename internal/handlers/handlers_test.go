package handlers

import (
	"Nexign_Career_Upgrade_24h/internal/config"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAPIJSON(t *testing.T) {

	stringRequest1 := "localhost" + ":" + config.Port + "/api/post/json"
	urlTestTestAPIJSON := `{"texts":[
    "Ощибка строка один",
    "Точная строка два",
    "зочем строка три"
	]}`
	ans := `[{"word":"Ощибка","s":["Ошибка"]},{"word":"зочем","s":["зачем","зочем"]}]`

	buffer := new(bytes.Buffer)
	buffer.WriteString(urlTestTestAPIJSON)

	req2, err := http.NewRequest(http.MethodPost, stringRequest1, buffer)
	req2.Header.Set("content-type", "application/json")
	if err != nil {
		fmt.Println(`err t2.TestApiJson!`)
		t.Fatalf("not req :%v", err)
	}
	rec2 := httptest.NewRecorder()
	GetAPIJSON(rec2, req2)
	t21 := rec2.Body

	if t21.String() == ans && rec2.Code == http.StatusOK {
		fmt.Println(`OK`)
	} else {
		fmt.Println(t21.String())
		fmt.Println(`is not -> `)
		fmt.Println(ans)
		fmt.Println(rec2.Code)
		fmt.Println(`err !!`)
	}

}
