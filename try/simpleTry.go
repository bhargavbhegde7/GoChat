package main

import(
		"encoding/json"
		"github.com/fatih/color"
)

type Request struct {
	Reqtag string `json:"reqtag"`
	Username string `json:"username"`
	Pubkey string `json:"pubkey"`
}

func main(){
	username := "bhegde"
	jsonString := `{"reqtag":"~&#signup#&~","username":"`+username+`","pubkey":"abcdef"}`

	color.Green(jsonString)

	request := Request{}
	err := json.Unmarshal([]byte(jsonString), &request)
	if err != nil{
		panic(err)
	}

	color.Yellow(request.Username)
}
