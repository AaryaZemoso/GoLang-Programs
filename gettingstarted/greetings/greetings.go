package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hello(s string) (string, error) {

	if s == "" {
		return "", errors.New("error : provide name")
	}

	message := fmt.Sprintf(randomFormat(), s)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi %v. Welcome",
		"Welcome %v.",
		"Nice to meet you, %v",
	}

	return formats[rand.Intn(len(formats))]

}
