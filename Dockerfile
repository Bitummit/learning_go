FROM golang:1.23.1


WORKDIR /app/

#Dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cd cmd/main && CGO_ENABLED=0 GOOS=linux go build -o /go_api

EXPOSE 8000

# TODO: mb run migrations

CMD ["/go_api"]