package main

import "fmt"
import "flag"
import "time"
import "github.com/op/go-nanomsg"

var verbose bool
var endProcessingCh chan bool
var processMsgCh chan []byte

func main() {
	ver := flag.Bool("Verbose", false, "Enable logging")
	flag.Parse()
	verbose = *ver
	endProcessingCh = make(chan bool)
	processMsgCh = make(chan []byte, 1000)
	go processMsg()
	go NanoSub()
	<-endProcessingCh
	fmt.Println("==========", cnt1)
}

var cnt1 int

func NanoSub() {
	sub, err := nanomsg.NewSubSocket()
	if err != nil {
		fmt.Println("Failed to open sub socket")
		return
	}
	ep, err := sub.Connect("ipc:///tmp/test.ipc")
	if err != nil {
		fmt.Println("Failed to connect to pub socket - ", ep)
		return
	}
	err = sub.Subscribe("")
	if err != nil {
		fmt.Println("Failed to subscribe to all topics")
		return
	}
	err = sub.SetRecvBuffer(1024 * 1024)
	if err != nil {
		fmt.Println("Failed to set recv buffer size")
		return
	}

	for {
		msg, err := sub.Recv(0)
		if err != nil {
			fmt.Println("Error in recv", err)
			continue
		}
		cnt1++
		processMsgCh <- msg
	}
	sub.Unsubscribe("")
}

func processMsg() {
	var startTime, endTime time.Time
	var msg []byte
	cnt := 0
	for {
		msg = <-processMsgCh
		cnt++
		if string(msg) == "Start" {
			startTime = time.Now()
			fmt.Println("Starting test:", startTime)
		} else if string(msg) == "End" {
			endTime = time.Now()
			fmt.Println("Ending test:", endTime)
			fmt.Println("Count:", cnt)
			endProcessingCh <- true
			break
		}
		if verbose == true {
			fmt.Printf("Message: %s\n", msg)
		}

	}
	fmt.Println("Time Duration:", endTime.Sub(startTime))
}
