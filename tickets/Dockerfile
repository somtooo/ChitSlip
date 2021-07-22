FROM golang:1.16.4

WORKDIR /go/src/github.com/BeatAllTech/ChitSlip/tickets
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

COPY . .


ENTRYPOINT CompileDaemon \
    --color="true" \
    --build="go build src/index.go" \
    --command="go run src/index.go"

