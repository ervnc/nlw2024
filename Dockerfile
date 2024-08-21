FROM golang:1.22.4-alpine

WORKDIR /nlw

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

WORKDIR /nlw/cmd/nlw

RUN go build -o /nlw/bin/nlw .

EXPOSE 8080
ENTRYPOINT [ "/nlw/bin/nlw" ]