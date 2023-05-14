FROM golang:1.20 as builder
RUN apt update && apt install git ca-certificates

WORKDIR /usr/src/app
COPY . .

ENV GO111MODULE=on

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ./main .

### Executable Image
FROM alpine

COPY --from=builder /usr/src/app/main .

RUN touch .env

EXPOSE 80

ENTRYPOINT ["/main"]