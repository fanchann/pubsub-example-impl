package services

import "fmt"

type ISubscriber interface {
	Notify(notif interface{})
}

type EmailSubscriber struct {
	Email string
}

func (e *EmailSubscriber) Notify(notif interface{}) {
	fmt.Printf("Email Notification sent to %s: %v \n", e.Email, notif)
}

type SMS struct {
	Number string
}

func (s *SMS) Notify(notif interface{}) {
	fmt.Printf("SMS sent to %s : %v \n", s.Number, notif)
}
