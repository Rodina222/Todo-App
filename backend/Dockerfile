FROM golang:1.20 AS builder

WORKDIR /app

COPY go.sum go.mod ./

RUN go mod download 

COPY . ./

RUN CGO_ENABLED=1  go build -o bin/app ./cmd/main.go


FROM alpine:latest  

WORKDIR /app

COPY --from=builder ./app/bin/app ./

EXPOSE 8096

ENV DATABASE_FILE=todoapp.db

CMD ["./app", "-db", "${DATABASE_FILE}", "-p", "8096"]
