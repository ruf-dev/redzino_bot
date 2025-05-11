package main

import (
	"github.com/sirupsen/logrus"

	"github.com/ruf-dev/redzino_bot/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		logrus.Fatal(err)
	}

	err = a.Start()
	if err != nil {
		logrus.Fatal(err)
	}
}
