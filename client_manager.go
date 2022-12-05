package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
	"websocketTest/types"
)


type ClientRoom struct {
	clientMap map[*Client]bool
	ChanEnter chan *Client
	ChanLeave chan *Client
	broadcast chan []byte		

	HostLastQueue *types.Queue
	BasicQueue *types.Queue
	IOQueue *types.Queue
}

func newClientRoom () *ClientRoom {
	return &ClientRoom{
		clientMap: make(map[*Client]bool, 5),
		broadcast: make(chan []byte),

		HostLastQueue: types.NewQueue(),
		BasicQueue: types.NewQueue(),		
		IOQueue: types.NewQueue(),	    
	}
}

func (w *ClientRoom) dataSet(tickString string, code uint32, dataQ *types.Queue) { 

	var header string 

	if (code == types.HOST_CODE) {
		header = "host, last "
	} else if (code == types.BASIC_CODE) {		
		header = "basic, cpu, mem "		
	} else if (code == types.NET_CODE) {
		header = "net, disk "		
	} 

	count := 1000
	for i := 0; i < count; i++ {

	var buf bytes.Buffer			  
	buf.WriteString(header)
	buf.WriteString(tickString)

	buf.WriteString(strconv.Itoa(i))	

	realData := types.RealData{}
	realData.Code = code				
	realData.Data = buf.String()	

	data, err := json.Marshal(realData)
	if err != nil {
		log.Printf("error: %s", err)
	}

	// dataQ.Push([]byte(data))
	dataQ.Push(data)
	}		
}

func (w *ClientRoom) dataProducer(code uint32, dataQ *types.Queue) { 

	ticker := time.NewTicker(1 * time.Second) //time.Millisecond
	var header string 

	if (code == types.HOST_CODE) {
		header = "host, last "
	} else if (code == types.BASIC_CODE) {		
		header = "basic, cpu, mem "		
	} else if (code == types.NET_CODE) {
		header = "net, disk "		
	} 

	for {
		select {            
			case tick := <- ticker.C: //1초마다 실행 
			  
			  count := 1000
			  for i := 0; i < count; i++ {

				var buf bytes.Buffer			  
				buf.WriteString(header)
				buf.WriteString(tick.String())

				buf.WriteString(strconv.Itoa(i))	

				realData := types.RealData{}
	            realData.Code = code				
				realData.Data = buf.String()	

				data, err := json.Marshal(realData)
				if err != nil {
					log.Printf("error: %s", err)
				}
                // dataQ.Push([]byte(data))
				dataQ.Push(data)
			  }						
		}				
	}
}

func (w *ClientRoom) dataPump(dQ *types.Queue) { //code uint32, 
	
	//우선 받은것 그대로 보내자. 
	// buf := types.RealData{}
	// buf.Code = code
		
	if (len(w.clientMap) == 0) {
		return
	}

	count := dQ.Count()
	for i := 0; i < count; i++ {
		//buf.Data += dQ.Pop()
		data := dQ.Pop().([]byte)
		for client := range w.clientMap {		
			// client.send <- dQ.Pop().([]byte)			
			// buf := dQ.Pop().(types.RealData);
			//buf := dQ.Pop().(string);
			// message, err := json.Marshal(buf)
			// if err != nil {
			// 	log.Printf("error: %s", err)
			// }
			client.send <- data
		}		
	}   		
}

//run 시에 ChanEnter, ChanLeave 를 초기화 한다. 
func (w *ClientRoom) run() {
	w.ChanEnter = make(chan *Client)
	w.ChanLeave = make(chan *Client)

	pingtTicker := time.NewTicker(1 * time.Second) //1초 Ticker 
	dataTicker := time.NewTicker(5 * time.Millisecond) 
	

	for {
		select {
		case client := <- w.ChanEnter: //클라이언트 접속시 알림 채널 
			fmt.Println("connect client(", client.id,")")
			w.clientMap[client] = true //클라이언트 맵에 True 표시 

		case client := <- w.ChanLeave: //클라이언트 접속 끊길시 알림 채널. 
			if _, ok := w.clientMap[client]; ok {
				delete(w.clientMap, client) //맵에서 제거 				
				fmt.Println("disconnect client(", client.id,")")
				close(client.send) //client 의 send 채널 close 
			}
			
		case message := <- w.broadcast:

			for client := range w.clientMap {		
				client.send <- message			
				// message = []byte(realdata) 
				// client.send <- message
				// // message, err = json.Marshal(buf)
				// // if err != nil {
				// // 	log.Printf("error: %s", err)
				// // }
			}		
			
		case <- dataTicker.C: 	//datatick :=<- dataTicker.C: 	
			//w.dataProducer(types.HOST_CODE, w.HostLastQueue)
			// w.dataProducer(types.BASIC_CODE, w.BasicQueue)
			// w.dataProducer(types.NET_CODE, w.IOQueue)
            //w.dataSet(datatick.String(), types.HOST_CODE, w.HostLastQueue)
		    w.dataPump(w.HostLastQueue)
			// w.dataPump(w.BasicQueue)
			// w.dataPump(w.IOQueue)
		    
					
		case tick := <- pingtTicker.C: //1초마다 실행 
			// w.dataProducer(types.HOST_CODE, w.HostLastQueue)
			// w.dataPump(w.HostLastQueue)
			

			// for client := range w.clientMap { //Client Map 전체 for 
			// 	client.send <- []byte(tick.String()) //tick 값 Send >> client.go 의 WritePump 에서 처리  
			// 	//client.send <- []byte("1,2,3,4,5,5,6,9") //tick 값 Send >> client.go 의 WritePump 에서 처리  
			// }
			// w.broadcast <- []byte(tick.String()) //tick 값 Send >> client.go 의 WritePump 에서 처리   
			fmt.Println(tick.String())
		}
	}
}

