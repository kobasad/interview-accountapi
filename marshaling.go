package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func marshal[T any](payload T) ([]byte, error) {
	log.Debug("Marshaling payload...")

	buf, err := json.Marshal(payload)

	if err != nil {
		log.WithField("error", err).Error("Failed to marshal payload request")
		return nil, err
	}

	log.WithFields(log.Fields{
		"buf": string(buf),
	}).Debug("Payload marshaled")

	return buf, nil
}

func unmarshal[T any](body []byte, payload T) error {
	log.Debug("Unmarshaling payload... ")

	err := json.Unmarshal(body, payload)

	if err != nil {
		log.WithField("error", err).Error("Failed to unmarshal the payload")
		return err
	}

	log.WithField("response", payload).Debug("Payload has been unmarshaled")

	return nil
}
