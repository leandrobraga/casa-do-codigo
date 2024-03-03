FROM golang:latest

WORKDIR /cdc
COPY . .

RUN go install github.com/cosmtrek/air@latest
# RUN go get -d -v ./...
# RUN go install -v ./...
RUN go mod tidy

RUN go build -o app ./cmd/api

CMD ["air"]