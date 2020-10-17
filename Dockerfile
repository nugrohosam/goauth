FROM golang:alpine

RUN apk update && apk upgrade

WORKDIR /app

COPY . .

RUN go get -v

RUN cp .env.example .env.yaml

RUN migrate -state=up -version=ver1

RUN go build -ldflags='[pattern=]args list'

EXPOSE 8080

CMD ["gosampleapi", "stage=dev"]