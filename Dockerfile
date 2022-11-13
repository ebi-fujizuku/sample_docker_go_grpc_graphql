FROM golang:1.18.2-bullseye

WORKDIR /article

COPY ./article /article

RUN go mod tidy

CMD ["go","run","./server/main.go"]
