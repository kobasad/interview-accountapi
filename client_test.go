package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	log "github.com/sirupsen/logrus"

	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const localServerBaseUrl = "http://localhost:8080"

// world stores the state of a particular scenario
var world struct {
	client         AccountApiClient
	createdAccount *AccountData
	fetchedAccount *AccountData
	mockServer     *httptest.Server
	err            error
}

func thereIsAClient() error {
	world.client = NewClient(localServerBaseUrl)
	return nil
}

func aNewAccountIsCreated() error {

	return assertActual(
		assert.NotEmpty, world.createdAccount,
		"Expected an account to be created", world.createdAccount,
	)
}

func createAccountIsCalled() error {

	var err error
	world.createdAccount, world.err = world.client.Create(aTestAccount())

	if world.err != nil {
		log.WithField("error", err).Error("Failed to create an account")
	}

	log.WithField("account", world.createdAccount).Debug("Account created")

	return nil
}

// Generates a test AccountData instance
func aTestAccount() *AccountData {
	version := int64(0)

	return &AccountData{
		Attributes:     someAccountAttributes(),
		ID:             uuid.New().String(),
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Version:        &version,
	}
}

func someAccountAttributes() *AccountAttributes {

	country := "GB"

	return &AccountAttributes{
		Country:      &country,
		BaseCurrency: "GBP",
		BankID:       "400300",
		BankIDCode:   "GBDSC",
		Bic:          "NWBKGB22",
		Name:         []string{"Samantha Holder"},
		UserDefinedData: []KeyValue{
			{
				Key:   "Some account related key",
				Value: "Some account related value",
			},
		},
		ValidationType:      "card",
		ReferenceMask:       "############",
		AcceptanceQualifier: "same_day",
	}
}

func deleteAccountIsCalled() error {

	err := world.client.Delete(world.createdAccount.ID, *world.createdAccount.Version)

	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"accountId": world.createdAccount.ID,
		}).Error("Failed to delete an account")
		return err
	}

	return nil
}

func fetchAccountIsCalled() error {

	var err error
	world.fetchedAccount, err = world.client.Fetch(world.createdAccount.ID)

	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"accountId": world.createdAccount.ID,
		}).Error("Failed to fetch an account")
		return err
	}

	return nil
}

func theAccountIsDeleted() error {

	var err error
	world.fetchedAccount, err = world.client.Fetch(world.createdAccount.ID)

	return assertExpectedAndActual(
		assert.Equal, ErrNotFound, err,
		"Expected account with ID %s to be deleted, but it's not", world.createdAccount.ID,
	)

}

func theExistingAccountIsReturned() error {

	return assertExpectedAndActual(
		assert.Equal, world.createdAccount, world.fetchedAccount,
		"Expected fetched account to be %v but was %v", world.createdAccount, world.fetchedAccount,
	)
}

func thereIsAnAccount() error {
	return createAccountIsCalled()
}

func errDuplicateAccountIsReturned() error {
	return assertExpectedAndActual(
		assert.Equal, ErrDuplicateAccount, world.err,
		"Expected ErrDuplicateAccount to be returned",
	)
}

func theSameAccountIsCreated() error {

	world.createdAccount, world.err = world.client.Create(world.createdAccount)

	return nil
}

func accountAPIWillReturnHTTPStatus(httpStatus int) error {

	world.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(httpStatus)
		w.Write([]byte(``))
	}))

	world.client = NewClient(world.mockServer.URL)

	return nil
}

func errRemoteServerIsReturned() error {
	return assertExpectedAndActual(
		assert.Equal, ErrRemoteServer, world.err,
		"Expected ErrRemoteServer but was %s", world.err)
}

func errUnknownIsReturned() error {
	return assertExpectedAndActual(
		assert.Equal, ErrUnknown, world.err,
		"Expected ErrUnknown to be returned",
	)
}

func httpCallWillResultInAnError() error {

	world.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		world.mockServer.CloseClientConnections()
	}))

	world.client = NewClient(world.mockServer.URL)

	return nil
}

func errNotFoundIsReturned() error {
	return assertExpectedAndActual(
		assert.Equal, ErrNotFound, world.err,
		"Expected ErrNotFound to be returned",
	)
}

func fetchAccountIsCalledForANonexistingAccount() error {
	world.fetchedAccount, world.err = world.client.Fetch(uuid.NewString())

	return nil
}

func deleteAccountIsCalledForANonexistingAccount() error {
	world.err = world.client.Delete(uuid.NewString(), 0)

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	//log.SetLevel(log.DebugLevel)
	log.SetLevel(log.FatalLevel)

	ctx.AfterScenario(func(sc *godog.Scenario, err error) {
		if world.mockServer != nil {
			world.mockServer.Close()
		}
	})

	ctx.Step(`^there is a client$`, thereIsAClient)
	ctx.Step(`^a new account is created$`, aNewAccountIsCreated)
	ctx.Step(`^\'create account\' is called$`, createAccountIsCalled)
	ctx.Step(`^\'delete account is called\'$`, deleteAccountIsCalled)
	ctx.Step(`^\'fetch account\' is called$`, fetchAccountIsCalled)
	ctx.Step(`^the account is deleted$`, theAccountIsDeleted)
	ctx.Step(`^the existing account is returned$`, theExistingAccountIsReturned)
	ctx.Step(`^there is an account$`, thereIsAnAccount)
	ctx.Step(`^ErrDuplicateAccount is returned$`, errDuplicateAccountIsReturned)
	ctx.Step(`^the same account is created$`, theSameAccountIsCreated)
	ctx.Step(`^Account API will return HTTP (\d+) status$`, accountAPIWillReturnHTTPStatus)
	ctx.Step(`^ErrRemoteServer is returned$`, errRemoteServerIsReturned)
	ctx.Step(`^ErrUnknown is returned$`, errUnknownIsReturned)
	ctx.Step(`^HTTP call will result in an error$`, httpCallWillResultInAnError)
	ctx.Step(`^ErrNotFound is returned$`, errNotFoundIsReturned)
	ctx.Step(`^\'fetch account\' is called for a non-existing account$`, fetchAccountIsCalledForANonexistingAccount)
	ctx.Step(`^\'delete account\' is called for a non-existing account$`, deleteAccountIsCalledForANonexistingAccount)
}

// Everything below is in the courtesy of https://github.com/cucumber/godog/blob/main/_examples/assert-godogs/godogs_test.go

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// assertActual is a helper function to allow the step function to call
// assertion functions where you want to compare an actual value to a
// predined state like nil, empty or true/false.
func assertActual(a actualAssertion, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, actual, msgAndArgs...)
	return t.err
}

type actualAssertion func(t assert.TestingT, actual interface{}, msgAndArgs ...interface{}) bool

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}
