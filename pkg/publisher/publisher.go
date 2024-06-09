package publisher

import (
	"fmt"
)

type Publisher struct{}

func New() *Publisher {
	return &Publisher{}
}

func (p *Publisher) Publish(data any) {
	fmt.Println(data)
}
