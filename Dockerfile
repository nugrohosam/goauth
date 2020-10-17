FROM golang:alpine

RUN apk update && apk upgrade

WORKDIR /app

COPY . .

RUN go get

RUN cp .env.example .env.yaml

RUN go build -x main.go

EXPOSE 8080

CMD main