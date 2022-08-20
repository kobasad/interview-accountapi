package main

import (
	"bytes"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Knows how to POST to Account API
func (client *ApiClient) post(buf []byte) (*http.Response, error) {

	url := client.BaseUrl + "/v1/organisation/accounts"

	log.WithFields(log.Fields{
		"url": url,
	}).Debug("Sending POST request to Account API")

	resp, err := client.HttpClient.Post(url, "application/json", bytes.NewReader(buf))

	if err != nil {
		log.WithField("error", err).Error("Failed to send request to Account API")
		return nil, err
	}

	log.Debug("Request to Account API sent")

	return resp, nil
}

// Knows how to GET from Account API
func (client *ApiClient) get(accountId string) (*http.Response, error) {

	url := fmt.Sprintf("%s/v1/organisation/accounts/%s", client.BaseUrl, accountId)

	log.WithFields(log.Fields{
		"url": url,
	}).Debug("Sending GET request to Account API")

	resp, err := client.HttpClient.Get(url)

	if err != nil {
		log.WithField("error", err).Error("Failed to send request to Account API")
		return nil, err
	}

	log.Debug("Request to Account API sent")

	return resp, nil
}

// Knows how to DELETE from Account API
func (client *ApiClient) delete(accountId string, version int64) (*http.Response, error) {

	// TODO remove duplication
	url := fmt.Sprintf("%s/v1/organisation/accounts/%s?version=%d", client.BaseUrl, accountId, version)

	log.WithFields(log.Fields{
		"url": url,
	}).Debug("Sending DELETE request to Account API")

	req, err := http.NewRequest("DELETE", url, nil)

	resp, err := client.HttpClient.Do(req)

	if err != nil {
		log.WithField("error", err).Error("Failed to send request to Account API")
		return nil, err
	}

	log.Debug("Request to Account API sent")

	return resp, nil
}
