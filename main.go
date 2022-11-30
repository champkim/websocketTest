package main

import (
	"log"
	"net/http"
)


func main() {
	clientRoom := newClientRoom()
	go clientRoom.run()

	var id uint64

	http.HandleFunc("/", serveHome)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		id ++
		client := NewClient(id, clientRoom, conn) //Client conn 을 가지는 Clinet 생성 -ReadPump, WritePump..set
		client.room.ChanEnter <- client //Clinet 의 주소값을 넘긴다. 
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}
