
### About

Simple banking app for the purpose of learning golang

APIs:

GET /greet : Simple helloworld response

GET /api/time : Current UTC timing, tz query parameter can be provided to return time in a specific region. For e.g. /api/time?tz=Asia/Singapore

GET /customers : Get all customers info

GET /customers/{customer_id} : Get customer info of a particular customerId

POST /account : Create an account

POST /transaction : Create a debit or credit transaction. Required params : transactionType(DEBIT|CREDIT), accountId(int), amount(float)


### Getting Started

1. Install necessary dependencies. For e.g. go, mysql, db schema

2. Ensure environment variable are set correctly
For e.g.
SERVER_ADDRESS=localhost;SERVER_PORT=8000;DB_USER=root;DB_PASSWORD=pwd;DB_NAME=banking

3. Run go app

### Potential Improvement

- Better code structure, shifting some cross-cutting concern packages such as errs and logger to a separate go module.
- Better logging
- Unit tests are not fully covered
- Authorization & Authentication
- Caching for GET requests




