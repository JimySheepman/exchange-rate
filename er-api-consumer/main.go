package main

import (
	"er-api-consumer/config"
	"er-api-consumer/task"

	"github.com/jasonlvhit/gocron"
)

func init() {
	config.LoadEnvirment()
}

func main() {
	gocron.Every(30).Seconds().Do(task.Task)

	<-gocron.Start()
}
