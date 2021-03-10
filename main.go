package main

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	// sockets
	/*server.OnConnect("connection", func(so socketio.Socket) {
		log.Println("New user connected! ")

		// Other events
	})*/

	// Http connections
	http.Handle("/socket.io", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on port : 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
