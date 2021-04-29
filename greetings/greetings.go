package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    messages_map := make(map[string]string)

    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        messages_map[name] = message
    }
    return messages_map, nil
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v",
        "Hail, %v Well met!",
    }

    return formats[rand.Intn(len(formats))]
}
