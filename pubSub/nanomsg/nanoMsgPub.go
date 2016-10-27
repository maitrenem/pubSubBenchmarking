package main

import "fmt"
import "flag"
import "time"
import "github.com/op/go-nanomsg"

var verbose bool

func main() {
	numOfIter := flag.Int("NumOfIter", 1, "Number of Iterations")
	numOfBytes := flag.Int("NumOfBytes", 100, "Num of Bytes")
	ver := flag.Bool("Verbose", false, "Enable logging")

	flag.Parse()
	verbose = *ver

	NanoMsg(*numOfIter, *numOfBytes)
}

func connect() (*nanomsg.PubSocket, error) {
	pub, err := nanomsg.NewPubSocket()
	if err != nil {
		fmt.Println("Failed to open pub socket")
		return nil, err
	}
	ep, err := pub.Bind("ipc:///tmp/test.ipc")
	if err != nil {
		fmt.Println("Failed to bind pub socket - ", ep)
		return nil, err
	}
	err = pub.SetSendBuffer(1024 * 1024)
	if err != nil {
		fmt.Println("Failed to set send buffer size")
		return nil, err
	}
	return pub, nil
}

func NanoMsg(numOfIter, numOfBytes int) error {
	pub, err := connect()
	if err != nil {
		fmt.Println("Connect error", err)
		return err
	}
	var msg string
	for i := 0; i < numOfBytes; i++ {
		msg += "A"
	}
	msg += "\n"
	time.Sleep(time.Duration(1) * time.Second)
	publish(pub, []byte("Start"))
	time.Sleep(time.Duration(1000000) * time.Nanosecond)
	for i := 0; i < numOfIter; i++ {
		publish(pub, []byte(msg))
		time.Sleep(time.Duration(1000000) * time.Nanosecond)
	}
	publish(pub, []byte("End"))
	return nil
}

func publish(pub *nanomsg.PubSocket, msg []byte) {
	_, rv := pub.Send(msg, 0)
	if rv != nil {
		fmt.Println("Send() rv: ", rv)
	}
}
