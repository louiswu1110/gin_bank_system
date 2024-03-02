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
- install docker and docker-compose
- command `docker-compose up`
- server will run on port 8080
- 

## Clean Env Run
- command `docker-compose down`
- command `docker-compose build --no-cache`
- command `docker-compose up`

## Question 1 Given the root of a binary tree, invert the tree, and return its root.(Don’t use recursion.)
- /cmd/question/questionOne.go

## To Be Optimized
- password salt for member and admin
- middleware for authentication
- auth flow(like jwt or session)
- transaction add sing
- admin permission design
- db migration
- env setting
