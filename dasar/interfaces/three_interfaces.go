package dasar

import (
	"fmt"
)

type email struct {
	isSubscribed bool
	body         string
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

func (e email) cost() int {
	lengthBody := len(e.body)
	if e.isSubscribed {
		return lengthBody * 2
	} else {
		return lengthBody * 5
	}
}

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf("'%s' | Subscribed", e.body)
	} else {
		return fmt.Sprintf("'%s' | Not Subscribed", e.body)
	}

}
