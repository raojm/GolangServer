package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}

func main() {
	// if err := db.Init(); err != nil {
	// 	panic(fmt.Sprintf("mysql init failed with %+v", err))
	// }

	http.HandleFunc("/socket", socketHandler)
	// http.HandleFunc("/", service.IndexHandler)
	// http.HandleFunc("/api/count", service.CounterHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
