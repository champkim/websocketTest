package main

import (
	"fmt"
	"time"
)

type ClientRoom struct {
	clientMap map[*Client]bool
	ChanEnter chan *Client
	ChanLeave chan *Client
	broadcast chan []byte
}

func newClientRoom () *ClientRoom {
	return &ClientRoom{
		clientMap: make(map[*Client]bool, 5),
		broadcast: make(chan []byte),
	}
}

//run 시에 ChanEnter, ChanLeave 를 초기화 한다. 
func (w *ClientRoom) run() {
	w.ChanEnter = make(chan *Client)
	w.ChanLeave = make(chan *Client)

	ticker := time.NewTicker(1 * time.Second) //1초 Ticker 

	for {
		select {
		case client := <- w.ChanEnter: //클라이언트 접속시 알림 채널 
			fmt.Println("connect client(", client,")")
			w.clientMap[client] = true //클라이언트 맵에 True 표시 

		case client := <- w.ChanLeave: //클라이언트 접속 끊길시 알림 채널. 
			if _, ok := w.clientMap[client]; ok {
				delete(w.clientMap, client) //맵에서 제거 				
				fmt.Println("disconnect client(", client,")")
				close(client.send) //client 의 send 채널 close 
			}

		case message := <- w.broadcast:
			for client := range w.clientMap {
				client.send <- message //
			}

		case tick := <- ticker.C: //1초마다 실행 
			for client := range w.clientMap { //Client Map 전체 for 
				client.send <- []byte(tick.String()) //tick 값 Send >> client.go 의 WritePump 에서 처리  
				//client.send <- []byte("1,2,3,4,5,5,6,9") //tick 값 Send >> client.go 의 WritePump 에서 처리  
			}
			fmt.Println(tick.String())
		}
	}
}

