FROM golang:latest as builder

RUN apt-get update && apt-get install -y ca-certificates tzdata

WORKDIR ./app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main .

FROM scratch

COPY --from=builder /go/app/main .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo/
