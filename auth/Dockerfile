FROM golang:alpine

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .

RUN cd ./cmd/webserver/ && go build -o main . && cp main ../../ && cd ../../

CMD ["./main"]