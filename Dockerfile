FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN cgo_enabled=0 GOOS=linux go build -o /donorbox-backend

EXPOSE 4000

CMD ["/donorbox-backend"]
