# Running

The easiest way would be to use the docker-compose. From the root directory:

```
docker compose up
```

Access the API at `http:localhost:8080`

# Implementation

I broke the implementation into layers

- API
- Controller
- Service
    - Receipt Service
    - Rule Service
- Repo

## API

The API layer creates a gin router and maps endpoints to the controller layer

## Controller

The controller handles the request and response tasks, namely serialization, deserialization, and status codes. It
delegates business logic to the service layer

## Service

The service layer handles logic related to the receipts. The `ReceiptService` processes new receipts and fetches the
points for existing receipts.

Points are stored with the receipt object, but it is done lazily. `RulesSerivce` is only called the first time the points are queried. After that, the points are fetched via the repository. 

In the `RuleService`, each rule is implemented as a struct with an `evaluate(*Receipt) int` method. The number of points earned on a receipt is equal to the sum of all `evaluate` methods.

## Repository

The repository offers an interface for getting receipts by id and storing new receipts. The `InMemoryReceiptRepository`
stores the data in a `map` that uses locking to achieve thread safety

## Tests
I wrote unit and integration tests for the service. All tests are appended with `_test`
