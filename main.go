package main

import (
	"fmt"
	"github.com/madhatter/tweetfeed/internal/config"
)

type configuration struct {
	hashtags string
	outfile  string
	loglevel int
	lastId   int
}

func main() {
	config := config.GetConfig()
	fmt.Printf("\n%s\n%d\n", config.Hashtags, config.LastId)

	config.SetHashtags("mesos pesos")
	fmt.Printf("\n%s\n%d\n", config.Hashtags, config.LastId)
}
