FROM golang:1.17

WORKDIR /wwgt-diary-server
COPY backend/go.mod .
RUN go mod download
COPY backend .

RUN go build -o main .
EXPOSE 8080

CMD ["./main"]