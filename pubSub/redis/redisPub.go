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

func publish(dbHdl redis.Conn, channel, value interface{}) {
	dbHdl.Do("PUBLISH", channel, value)
}

func RedisPub(numOfItr, numOfBytes int) {
	c, err := dial()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	var msg string
	for i := 0; i < numOfBytes; i++ {
		msg += "A"
	}
	publish(c, "", "Start")
	for i := 0; i < numOfItr; i++ {
		publish(c, "", msg)
	}
	publish(c, "", "End")
}

func main() {
	numOfItr := flag.Int("NumOfIter", 1, "Number of Iterations")
	numOfBytes := flag.Int("NumOfBytes", 100, "Num of Bytes")
	ver := flag.Bool("Verbose", false, "Enable logging")
	flag.Parse()
	verbose = *ver
	RedisPub(*numOfItr, *numOfBytes)
}
