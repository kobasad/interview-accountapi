package main

import "errors"

var ErrNotFound = errors.New("Not Found")
var ErrDuplicateAccount = errors.New("Account with the given ID already exists")
var ErrRemoteServer = errors.New("Remote Server Error")
var ErrUnknown = errors.New("Unknown Error")

type CreateAccountRequest struct {
	Data *AccountData `json:"data,omitempty"`
}

type CreateAccountResponse struct {
	Data *AccountData `json:"data,omitempty"`
}

type GetAccountResponse struct {
	Data *AccountData `json:"data,omitempty"`
}

// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type KeyValue struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string    `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool      `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string     `json:"account_number,omitempty"`
	AlternativeNames        []string   `json:"alternative_names,omitempty"`
	BankID                  string     `json:"bank_id,omitempty"`
	BankIDCode              string     `json:"bank_id_code,omitempty"`
	BaseCurrency            string     `json:"base_currency,omitempty"`
	Bic                     string     `json:"bic,omitempty"`
	Country                 *string    `json:"country,omitempty"`
	Iban                    string     `json:"iban,omitempty"`
	JointAccount            *bool      `json:"joint_account,omitempty"`
	Name                    []string   `json:"name,omitempty"`
	SecondaryIdentification string     `json:"secondary_identification,omitempty"`
	Status                  *string    `json:"status,omitempty"`
	Switched                *bool      `json:"switched,omitempty"`
	UserDefinedData         []KeyValue `json:"user_defined_data,omitempty"`
	ValidationType          string     `json:"validation_type,omitempty"`
	ReferenceMask           string     `json:"reference_mask,omitempty"`
	AcceptanceQualifier     string     `json:"acceptance_qualifier,omitempty"`
}
