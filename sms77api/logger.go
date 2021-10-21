package sms77api

import (
	"fmt"
	"log"
)

type logger struct {
}

func NewDefaultLogger() *logger {
	return &logger{}
}

func (l *logger) Debug(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("DEBUG: %s", msg)
}

func (l *logger) Error(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("ERROR: %s", msg)
}
