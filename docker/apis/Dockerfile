FROM golang:1.19-alpine
ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor
WORKDIR /gin_bank
COPY ../.. .
RUN go build -o main ./cmd/server
EXPOSE 8080
CMD ["./main"]