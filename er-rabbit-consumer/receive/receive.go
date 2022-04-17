package receive

import (
	"encoding/json"
	"log"

	"er-rabbit-consumer/db"
	"er-rabbit-consumer/model"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ReceiveTRY() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TRY", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		db, err := db.ConnectDB()
		if err != nil {
			log.Fatal("Cannot connect to database")
		}

		for d := range msgs {
			data := model.Currencies{}
			json.Unmarshal(d.Body, &data)

			insertDynStmt := `insert into currencies(base_code, target_code, conversion_rate, created_at) values($1, $2,$3,$4)`
			_, err = db.Exec(insertDynStmt, data.BaseCode, data.TargetCode, data.ConversionRate, data.CreatedAt)
			if err != nil {
				log.Fatal("Cannot connect to database")
			}

			log.Printf("Received a message: %s", d.Body)
		}

	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func ReceiveUSD() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"USD", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		db, err := db.ConnectDB()
		if err != nil {
			log.Fatal("Cannot connect to database")
		}

		for d := range msgs {
			data := model.Currencies{}
			json.Unmarshal(d.Body, &data)

			insertDynStmt := `insert into currencies(base_code, target_code, conversion_rate, created_at) values($1, $2,$3,$4)`
			_, err = db.Exec(insertDynStmt, data.BaseCode, data.TargetCode, data.ConversionRate, data.CreatedAt)
			if err != nil {
				log.Fatal("Cannot connect to database")
			}

			log.Printf("Received a message: %s", d.Body)
		}

	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func ReceiveEUR() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"EUR", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		db, err := db.ConnectDB()
		if err != nil {
			log.Fatal("Cannot connect to database")
		}

		for d := range msgs {
			data := model.Currencies{}
			json.Unmarshal(d.Body, &data)

			insertDynStmt := `insert into currencies(base_code, target_code, conversion_rate, created_at) values($1, $2,$3,$4)`
			_, err = db.Exec(insertDynStmt, data.BaseCode, data.TargetCode, data.ConversionRate, data.CreatedAt)
			if err != nil {
				log.Fatal("Cannot connect to database")
			}

			log.Printf("Received a message: %s", d.Body)
		}

	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
