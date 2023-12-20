package services

type IPublisher interface {
	AddSubscriber(s ISubscriber)
	RemoveSubcriber(s ISubscriber)
	Publish(msg interface{})
}

type Publisher struct {
	Subscriber []ISubscriber
}

func (p *Publisher) AddSubscriber(s ISubscriber) {
	p.Subscriber = append(p.Subscriber, s)
}

func (p *Publisher) RemoveSubcriber(s ISubscriber) {
	for i, subscriber := range p.Subscriber {
		if subscriber == s {
			p.Subscriber = append(p.Subscriber[:i], p.Subscriber[i+1:]...)
			break
		}
	}
}

func (p *Publisher) Publish(msg interface{}) {
	for _, subscriber := range p.Subscriber {
		subscriber.Notify(msg)
	}
}
