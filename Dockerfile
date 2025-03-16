FROM golang:1.24.1

RUN go version
ENV GOPATH=/
COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client
RUN chmod +x waiting.sh

RUN go mod download
RUN go build -o restapigarbage ./cmd/main.go

CMD ["./app"]
