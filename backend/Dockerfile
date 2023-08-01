FROM golang:1.20 AS builder

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0  go build -o app cmd/main.go


FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder ./app ./

EXPOSE 8096

CMD ["./app", "-db", "${DATABASE_FILE}"]