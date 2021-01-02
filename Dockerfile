FROM golang:alpine

RUN apk update && apk upgrade

WORKDIR /app

COPY . .

RUN go get -v

RUN cp .env.yaml.example .env.yaml

RUN go build -ldflags='[pattern=]args list'

EXPOSE 8080

CMD ["app", "stage=dev"]