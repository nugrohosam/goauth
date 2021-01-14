FROM golang:alpine

RUN apk update && apk upgrade

WORKDIR /app

COPY . .

RUN cp .env.yaml.example .env.yaml

RUN make

EXPOSE 8080

CMD ["app", "stage=dev"]