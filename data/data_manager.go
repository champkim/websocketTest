package data

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"
	"time"
	"websocketTest/types"
)

//var requestChan chan chan make(chan chan *types.Queue)

type DataQueue struct {
	RequestChan chan chan *types.Queue
	dataQueue *types.Queue	
	targetQueue *types.Queue
}

func newDataQueue () *DataQueue {
	return &DataQueue{
		RequestChan: make(chan chan *types.Queue),
		dataQueue: types.NewQueue(),
		targetQueue: types.NewQueue(),
	}
}

func (d *DataQueue) SetTargetQueue(targetQ **types.Queue) {
	d.targetQueue = *targetQ
}

func (d *DataQueue) DataProducer(code uint32) { 

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
			case responseChan := <-d.RequestChan:
				types.ChangeQueue(&d.dataQueue, &d.targetQueue)
				responseChan <- d.targetQueue


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
				d.dataQueue.Push(data)				
			  }						
		}				
	}
}



// func (w *DataRoom) dataPump() {

// 	dataTicker := time.NewTicker(500*time.Microsecond)
// 	for {
// 		select {

// 		case message := <- w.broadcast:

// 			for client := range w.clientMap {

// 				var realdata string = "client[" + strconv.Itoa(int(client.id)) + "]";
// 				realdata += " datakey["+ strconv.Itoa(int(client.dataKey))+"] "

// 				if client.dataKey.IsSet(types.HOST_KEY) {
// 					realdata += ",host"
// 				}
// 				if client.dataKey.IsSet(types.LASTPERF_KEY) {
// 					realdata += ",last"
// 				}
// 				if client.dataKey.IsSet(types.BASIC_KEY) {
// 					realdata += ",basic"
// 				}
// 				if client.dataKey.IsSet(types.CPU_KEY) {
// 					realdata += ",cpu"
// 				}
// 				if client.dataKey.IsSet(types.MEM_KEY) {
// 					realdata += ",mem"
// 				}
// 				if client.dataKey.IsSet(types.NET_KEY) {
// 					realdata += ",net"
// 				}
// 				if client.dataKey.IsSet(types.DISK_KEY) {
// 					realdata += ",disk"
// 				}
// 				realdata += string(message)

// 				message = []byte(realdata)
// 				client.send <- message
// 				// message, err = json.Marshal(buQueue
// 				// if err != nil {
// 				// 	log.Printf("error: %s", err)
// 				// }
// 			}

// 		case <- dataTicker.C: //1초마다 실행
// 			for client := range w.clientMap { //Client Map 전체 for
// 				client.send <- []byte(tick.String()) //tick 값 Send >> client.go 의 WritePump 에서 처리
// 				//client.send <- []byte("1,2,3,4,5,5,6,9") //tick 값 Send >> client.go 의 WritePump 에서 처리
// 			}
// 			//w.broadcast <- []byte(tick.String()) //tick 값 Send >> client.go 의 WritePump 에서 처리
// 			fmt.Println(tick.String())
// 		}
// 	}

// }