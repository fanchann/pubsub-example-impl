package main

import "github.com/fanchann/belajar-pubsub/services"

func main() {
	publisher := services.Publisher{}

	emailSubscriber := &services.EmailSubscriber{Email: "test@gmail.com"}
	smsSubscriber := &services.SMS{Number: "6666666666s6"}

	publisher.AddSubscriber(emailSubscriber)
	publisher.AddSubscriber(smsSubscriber)

	publisher.RemoveSubcriber(smsSubscriber)

	publisher.Publish("Hi")
}
