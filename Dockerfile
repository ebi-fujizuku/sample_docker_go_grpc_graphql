FROM golang:1.18.2-bullseye

WORKDIR /article

COPY ./article /article

RUN apt-get update
# RUN go install github.com/joho/godotenv/cmd/godotenv@latest
RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest
RUN go mod tidy

# CMD ["go","run","./server/main.go"]
