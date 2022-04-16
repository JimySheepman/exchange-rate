package task

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"er-api-consumer/model"

	"github.com/streadway/amqp"
)

const (
	MAIN_PATH = "https://v6.exchangerate-api.com/v6/"
	SUB_PATH  = "/latest/"
)

var queryPaths = [3]string{"TRY", "USD", "EUR"}

func Task() {
	var rates model.Rates
	_ = rates

	for _, path := range queryPaths {
		fecthValue := GetValueFromAPI(path)
		Send(fecthValue)
	}
}

func GetValueFromAPI(path string) []byte {
	res, err := http.Get(MAIN_PATH + os.Getenv("API_KEY") + SUB_PATH + path)
	if err != nil {
		log.Fatal(err)
	}

	req, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Could not complete read from request body")
	}

	/* 	err = json.Unmarshal(req, &rates)
	   	if err != nil {
	   		log.Fatal("Could not complete unmarshal body")
	   	}
	*/
	//fmt.Printf("%s", req)
	return req
}

func Send(fecthValue []byte) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        fecthValue,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", fecthValue)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
