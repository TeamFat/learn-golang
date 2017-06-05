/**
 * datetime: 2017-06-05
 * Go Json操作 - 解码操作
 */
package main

import (
	"encoding/json"
	"fmt"
)

type Game struct {
	Id   int    `json:"game_id"`
	Name string `json:"game_name"`
}

func main() {

	game := Game{Name: "传奇霸业", Id: 100}
	//编码
	jsonStr, _ := json.Marshal(game)
	fmt.Printf("json encode %s\n", jsonStr)
	//解码
	var v map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &v)
	fmt.Println(v["game_id"])
	fmt.Println(v["game_name"])
}

//
// json encode {"game_id":100,"game_name":"传奇霸业"}
// 100
// 传奇霸业
//
