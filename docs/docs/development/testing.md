# Testing

In order to ensure the quality of our codebase, we have implemented comprehensive testing strategies for both the frontend and backend components of the application.

[Test-Driven Development (TDD)](https://martinfowler.com/bliki/TestDrivenDevelopment.html) is recommended as a best practice, although it is not mandatory or enforced. This approach encourages writing tests before the actual implementation, leading to more reliable and maintainable code.

## Writing Tests

Tests follow the [Arrange-Act-Assert](https://automationpanda.com/2020/07/07/arrange-act-assert-a-pattern-for-writing-good-tests/) (AAA) pattern to enhance readability and maintainability. This pattern involves three steps:

**Arrange**: Set up the necessary preconditions and inputs. This may include creating objects, setting initial values, or configuring mocks implementations.
**Act**: Execute the logic being tested, such as calling a function or method.
**Assert**: Verify that the outcome is as expected.

## Frontend

The frontend test suite consists of:

- [MockServiceWorker (MSW)](https://mswjs.io/) for API mocking
- [Vitest](https://vitest.dev/) as the test runner
- [Vue Test Utils](https://test-utils.vuejs.org/) for component testing

### Running Frontend Tests

To run the frontend tests, use one of the following commands:

```bash
# from root dir
/> make test-fe

# from frontend dir
/api/frontend/> pnpm test # runs test

/api/frontend/> pnpm coverage # runs coverage
```

## Backend

The backend test suite consists of:

- [Testify](https://github.com/stretchr/testify) for assertions and mocking
- [Mockery](https://github.com/vektra/mockery) for generating mocks

### Running Backend Tests

To run the backend tests, use the following command:

```bash
/> make test-be
```

## End-to-End (e2e) Testing

End to End testing is implemented using [Playwright](https://playwright.dev/) with the [playwright-go](https://github.com/playwright-community/playwright-go) library.

### Writing e2e Tests

End to End tests SHOULD adhere to [Behavior-Driven Development](https://en.wikipedia.org/wiki/Behavior-driven_development) practices, which emphasize focusing on the User perspective and the associated scenario a User encounters within the system.

Behavior-Driven Development Tests should follow the Given-When-Then format:

- Given: the initial context at the beginning of the scenario, in one or more clauses;
- When: the event that triggers the scenario;
- Then: the expected outcome, in one or more clauses.

In order to develop e2e tests, the following development flow is recommended:

- Open 2 terminal sessions

    ```bash
    # Terminal 1
    make compose
    ```

- Write the associated Test
- Run the test locally

    ```bash
    # Terminal 2
    make e2e
    ```

### Running e2e Tests

Run with docker compose to handle spinning up the application + database + e2e tests within a container

```bash
make compose-e2e
```

Run locally (**assuming the app + database are already running**)

```bash
make e2e
```
