FROM golang:1.17.6-alpine

WORKDIR /app

COPY ./go_app/go.mod ./
COPY ./go_app/go.sum ./

RUN go mod download

COPY ./go_app/*.go ./

RUN go build -o /go_app

CMD ["/go_app"]

