package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "meleniawan").Info("Hello World")

	logger.WithField("username", "yoga").
		WithField("name", "Yoga Meleniawan").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "yoga",
		"name":     "Yoga Meleniawan",
	}).Info("Hello World")
}
