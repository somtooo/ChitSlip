FROM golang:1.16.4

WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon


COPY . .

CMD [ "CompileDaemon", "-command=\"go run index.go\"" ]