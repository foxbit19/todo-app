FROM golang:1.19-alpine

WORKDIR /usr/server

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/server

EXPOSE 8080
CMD ["server"]