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
/> make test-fe

/api/frontend/> pnpm test

/api/frontend/> pnpm coverage
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

## End-to-End (E2E) Testing

<!-- TODO: WIP -->