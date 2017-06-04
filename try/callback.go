package main

import (
    "fmt"
    "encoding/json"
)

type Data struct {
    Count string `json:"count"`
	}



func main() {
    s := `{ "count": "3" }`
    data := &Data{

    }
    err := json.Unmarshal([]byte(s), data)
    fmt.Println(err)
    fmt.Println(data.Count)
    // s2, _ := json.Marshal(data)
    // fmt.Println(string(s2))
    // data.Count = "2"
    // s3, _ := json.Marshal(data)
    // fmt.Println(string(s3))
}
