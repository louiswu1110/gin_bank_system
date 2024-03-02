# For meepshop project
## Default Data
- admin account
  - username: admin
  - id: 3
- member
  - username: ming
  - id: 1
  - username: bear
  - id: 2

## How to run
- install `docker` and `docker-compose`
- command `docker-compose up`
- server will run on port `8080`

## How to Call API
- use `postman` or `curl` to call API
- API document in `meepshop.postmanv2.1_collection.json`

## Clean Environment
- command `docker-compose down`
- command `docker-compose build --no-cache`

## Question 1 Given the root of a binary tree, invert the tree, and return its root.(Don’t use recursion.)
- /cmd/question/questionOne.go

## To Be Optimized
- password salt for member and admin
- middleware for authentication
- auth flow(like jwt or session)
- transaction add sing
- admin permission design
- DB migration
- env setting
- full unit test
- statistics for transaction (daily, weekly, monthly, yearly)
- CI/CD setting(github action, jenkins, etc.)