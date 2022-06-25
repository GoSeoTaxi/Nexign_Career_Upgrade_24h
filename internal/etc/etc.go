package etc

type AnsJ struct {
	Word string   `json:"word"`
	S    []string `json:"s"`
}

type TJSONIn []struct {
	code int      `json:"code"`
	pos  int      `json:"pos"`
	row  int      `json:"row"`
	col  int      `json:"col"`
	len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}
