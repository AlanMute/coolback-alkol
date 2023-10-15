FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go build -o eptanit ./cmd/main.go

EXPOSE 8080

CMD ["./eptanit"]