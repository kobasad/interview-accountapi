FROM golang:1.18

COPY . /app

WORKDIR /app

RUN go install github.com/cucumber/godog/cmd/godog@latest

RUN go build -v

RUN go install

CMD ["godog", "run"]