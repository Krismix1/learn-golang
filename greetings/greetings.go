package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	format := randomGreet()
	message := fmt.Sprintf(format, name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Greet(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomGreet() string {
	formats := []string{"Hi, %v. Welcome!", "Great to see you, %s!", "Hail, %v. Well met!"}
	randomIndex := rand.Intn(len(formats))
	return formats[randomIndex]
}
