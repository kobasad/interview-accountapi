Feature: Account API client

    As a developer
    I want an account API client
    So that I could use it without thinking too much about the integration details

    Background:
        Given there is a client

    Scenario: create new account
        When 'create account' is called
        Then a new account is created

    Scenario: create a duplicate
        Given there is an account
        When the same account is created
        Then ErrDuplicateAccount is returned

    Scenario Outline: API failure when trying to create a new account
        Given Account API will return HTTP <status> status
        When 'create account' is called
        Then ErrRemoteServer is returned

        Examples:
            | status |
            | 500    |
            | 501    |
            | 502    |
            | 503    |
            | 504    |
            | 505    |
            | 506    |
            | 507    |
            | 508    |
            | 509    |


    Scenario: network or hardware failure when trying to create a new account
        Given HTTP call will result in an error
        When 'create account' is called
        Then ErrUnknown is returned
    # TODO: request and result accounts specified as tables with all supported attributes

    Scenario: fetch an account
        Given there is an account
        When 'fetch account' is called
        Then the existing account is returned

    Scenario: fetch a non-existing account
        When 'fetch account' is called for a non-existing account
        Then ErrNotFound is returned

    Scenario Outline: API failure when trying to fetch an account
        Given Account API will return HTTP <status> status
        When 'fetch account' is called for a non-existing account
        Then ErrRemoteServer is returned

        Examples:
            | status |
            | 500    |
            | 501    |
            | 502    |
            | 503    |
            | 504    |
            | 505    |
            | 506    |
            | 507    |
            | 508    |
            | 509    |

    Scenario: network or hardware failure when fetching an account
        Given HTTP call will result in an error
        When 'fetch account' is called for a non-existing account
        Then ErrUnknown is returned

    Scenario: delete an account
        Given there is an account
        When 'delete account is called'
        Then the account is deleted

    Scenario: delete a non-existing account
        When 'delete account' is called for a non-existing account
        Then ErrNotFound is returned

    Scenario Outline: API failure when trying to delete an account
        Given Account API will return HTTP <status> status
        When 'delete account' is called for a non-existing account
        Then ErrRemoteServer is returned

        Examples:
            | status |
            | 500    |
            | 501    |
            | 502    |
            | 503    |
            | 504    |
            | 505    |
            | 506    |
            | 507    |
            | 508    |
            | 509    |

    Scenario: network or hardware failure when deleting an account
        Given HTTP call will result in an error
        When 'delete account' is called for a non-existing account
        Then ErrUnknown is returned
