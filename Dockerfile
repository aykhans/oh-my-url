FROM golang:1.21.6-alpine AS builder

WORKDIR /ohmyurl

COPY go.mod go.sum ./
RUN go mod download
COPY ./app ./app

RUN go build -o ./ohmyurl ./app/main.go

FROM scratch

COPY --from=builder /ohmyurl/app/templates/ /app/templates/
COPY --from=builder /ohmyurl/ohmyurl /ohmyurl

ENTRYPOINT ["/ohmyurl"]
