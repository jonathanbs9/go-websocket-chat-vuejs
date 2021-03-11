package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	// sockets
	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		so.Join("chat_room")
		fmt.Println("Nuevo usuario conectado! ")
		return nil
	})

	server.OnEvent("/", "chat message", func(so socketio.Conn, msg string) {
		log.Println(msg)
		server.BroadcastToRoom("chat message", "chat_room", msg)
	})


	go server.Serve()
	defer server.Close()

	// Http connections
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on port : 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))


}
