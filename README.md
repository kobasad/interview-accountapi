# Form3 Take Home Exercise

This is a take home exercise for Form3 implemented by me, Aleksandr Kobyshcha. This is my first experience with Go, I didn't know the language before starting this assignment.
The solution provides a module that can be used as a Form3 Account API client. The client was developed using BDD approach with help of `godog`, Go implementation of `Cucumber`. Corresponding feature file containing Gherkin scenarios is located in `features` directory.

# Executing Tests

Tests are executed using `docker-compose`, the docker image containing the solution is built during `docker-compose` startup, so building an image separately is not required. Tests can be executed in two ways:
- `docker-compose run test`: this will build the `test` container, execute the tests, and exit docker-compose
- `docker-compose up`: this will build the `test` container, execute the tests, and keep docker-compose running.

# Logging

Detailed logs can be enabled by setting `LOG_LEVEL` environment variable to `DEBUG`.

# Potential Improvements
- Test coverage could be improved by supplying test data via tables instead of code-defined test data
- Marshaling and read/write error handling should be done similarly to other errors: return custom error constant for each case
- In order to be production-ready, the solution should include monitoring capabilities and other quality attributes
- `http.Client` could be reused among all instances of `ApiClient` in order to optimize performance
- async client?