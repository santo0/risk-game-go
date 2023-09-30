package main

import (
	"fmt"
	"net/http"

	engine "github.com/santo0/risk-game-go/internal/engine"
	model "github.com/santo0/risk-game-go/internal/model"
)

const MAP_FILE string = "/home/user/Documents/hacknights_29092023/risk-game-go/assets/battlefield_map_1.txt"

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	players := []model.Player{
		{PlayerName: "joe", PlayerId: 1},
		{PlayerName: "bill", PlayerId: 2},
		{PlayerName: "mike", PlayerId: 3},
	}
	btf := engine.CreateBattlefieldMap(MAP_FILE, engine.DiceBattleConfiguration{})
	engine.CreateGame(players, btf)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
