FROM golang:1.17

WORKDIR /wwgt-diary-server
COPY go-diaries/go.mod go-diaries/go.sum ./
RUN go mod download
COPY go-diaries .

RUN go build -o main .
EXPOSE 8080

CMD ["./main"]