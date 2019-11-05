package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		return
	}

	logrus.SetLevel(logrus.DebugLevel)
}
