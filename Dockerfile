FROM golang:latest AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# RUN 用多個RUN去寫的話，會產生新的image
RUN go mod tidy && \
    go build -o ginChat


# CMD sleep 15 && ./ginChat
CMD ["./wait-for-it.sh", "mysql:3306", "--", "./ginChat"]
