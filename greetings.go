package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func Hello(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", errors.New("El nombre no puede estar vacío")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string, len(names))
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[strings.TrimSpace(name)] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hola, %s, Bienvenido!",
		"Saludos, %s!",
		"Bienvenido, %s!",
		"Qué tal, %s!",
		"Un placer conocerte, %s!",
		"Espero que tengas un gran día, %s!",
	}
	return formats[rand.Intn(len(formats))]
}
