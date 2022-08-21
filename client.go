package main

import (
	"errors"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Account API client interface
type AccountApiClient interface {
	// Creates a new Account
	Create(account *AccountData) (*AccountData, error)

	// Fetches an existing Account by its ID
	Fetch(accountId string) (*AccountData, error)

	// Deletes an existing Account by its ID
	Delete(accountId string, version int64) error
}

// Generic API client structure containing configuration data and the underlying HTTP client
type ApiClient struct {
	BaseUrl    string
	HttpClient *http.Client
}

// Factory method to instantiate an Account API client
func NewClient(baseUrl string) AccountApiClient {

	return &ApiClient{baseUrl, &http.Client{}}
}

func (client *ApiClient) Create(account *AccountData) (*AccountData, error) {

	request := CreateAccountRequest{
		Data: account,
	}

	buf, err := marshal(&request)

	if err != nil {
		return nil, err
	}

	resp, err := client.post(buf)

	if err != nil {
		return nil, ErrUnknown
	}

	body, err := read(resp)

	if err != nil {
		return nil, err
	}

	if err = toError(resp); err != nil {
		log.WithFields(log.Fields{
			"statusCode": resp.StatusCode,
			"body":       string(body),
			"error":      err,
		}).Error("Request to Account API failed")
		return nil, err
	}

	result := &CreateAccountResponse{}

	err = unmarshal(body, result)

	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (client *ApiClient) Fetch(accountId string) (*AccountData, error) {

	resp, err := client.get(accountId)

	if err != nil {
		return nil, ErrUnknown
	}

	body, err := read(resp)

	if err != nil {
		return nil, err
	}

	if err = toError(resp); err != nil {
		log.WithFields(log.Fields{
			"statusCode": resp.StatusCode,
			"body":       string(body),
			"error":      err,
		}).Error("Request to Account API failed")

		return nil, err
	}

	result := &GetAccountResponse{}

	err = unmarshal(body, result)

	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (client *ApiClient) Delete(accountId string, version int64) error {

	resp, err := client.delete(accountId, version)

	if err != nil {
		return ErrUnknown
	}

	if err = toError(resp); err != nil {
		log.WithField("statusCode", resp.StatusCode).Error("Request to Account API failed")
		return err
	}

	return nil
}

// Response reader
func read(resp *http.Response) ([]byte, error) {
	log.Debug("Reading the response... ")

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.WithField("error", err).Error("Failed to read the response")
		return nil, err
	}

	log.WithField("body", string(body)).Debug("Response has been read")

	return body, nil
}

var okStatuses = map[int]int{
	http.StatusOK:        1,
	http.StatusAccepted:  1,
	http.StatusNoContent: 1,
	http.StatusCreated:   1,
}

// Converts execution errors into meaningful client errors
func toError(resp *http.Response) error {
	switch statusCode := resp.StatusCode; {
	case okStatuses[statusCode] == 1:
		return nil
	case statusCode == http.StatusNotFound:
		return ErrNotFound
	case statusCode == http.StatusConflict:
		return ErrDuplicateAccount
	case statusCode >= 500:
		return ErrRemoteServer
	}

	return errors.New("Request to Account API failed")
}

func main() {}
