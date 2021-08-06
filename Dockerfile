FROM golang:alpine as builder

ENV GO111MODULE=on

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        && apk add curl \
        ca-certificates openssl \
        cp giag4.pem /usr/local/share/ca-certificates/giag4.pem \
        && update-ca-certificates 2>/dev/null || true

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

FROM scratch

COPY --from=builder /app/bin/main .

CMD ["./main"]