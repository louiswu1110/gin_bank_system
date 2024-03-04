# For gin project
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
- API document in `gin.postmanv2.1_collection.json`

## Clean Environment
- command `docker-compose down`
- command `docker-compose build --no-cache`

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
