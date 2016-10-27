package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var verbose bool

func dial() (redis.Conn, error) {
	retryCount := 0
	ticker := time.NewTicker(2 * time.Second)
	for _ = range ticker.C {
		retryCount += 1
		dbHdl, err := redis.Dial("tcp", ":6379")
		if err != nil {
			if retryCount%100 == 0 {
				fmt.Println("Failed to dail out to Redis server. Retrying connection. Num of retries = ", retryCount)
			}
		} else {
			return dbHdl, nil
		}
	}
	err := errors.New("Error opening db handler")
	return nil, err
}

var startTime, endTime time.Time

func RedisSub() {
	c, err := dial()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	dbHdl := redis.PubSubConn{Conn: c}
	defer dbHdl.Unsubscribe("")

	dbHdl.Subscribe("")
	//dbHdl.PSubscribe("p*")
	cnt := 0
	for {
		switch n := dbHdl.Receive().(type) {
		case redis.Message:
			cnt++
			if string(n.Data) == "Start" {
				startTime = time.Now()
				fmt.Println("Starting test:", startTime)
			} else if string(n.Data) == "End" {
				endTime = time.Now()
				fmt.Println("Ending test:", endTime)
				fmt.Println("Count:", cnt)
				return
			}
			if verbose == true {
				fmt.Printf("Message: %s %s\n", n.Channel, n.Data)
			}
		//case redis.PMessage:
		//	fmt.Printf("PMessage: %s %s %s\n", n.Pattern, n.Channel, n.Data)
		case redis.Subscription:
			fmt.Printf("Subscription: %s %s %d\n", n.Kind, n.Channel, n.Count)
			if n.Count == 0 {
				return
			}
		case error:
			fmt.Printf("error: %v\n", n)
			return
		}
	}
	//dbHdl.Unsubscribe("")
	//dbHdl.PUnsubscribe()
}

func main() {
	ver := flag.Bool("Verbose", false, "Enable logging")
	flag.Parse()
	verbose = *ver
	RedisSub()
	fmt.Println("Time Duration:", endTime.Sub(startTime))
}
