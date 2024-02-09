package main

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"
)

var address = ":4321"

func LengthServerResponse(ws *websocket.Conn) {
	var msg string

	for {
		websocket.Message.Receive(ws, &msg)
		fmt.Println("Got message", msg)
		length := len(msg)
		if err := websocket.Message.Send(ws, strconv.FormatInt(int64(length), 10)); err != nil {
			fmt.Println("Cannot send messageg length")
			break
		}
	}
}

func websocketListen() {
	http.Handle("/length", websocket.Handler(LengthServerResponse))
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	websocketListen()
}
