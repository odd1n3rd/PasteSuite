FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pingApp ./cmd

FROM alpine:latest

RUN addgroup -S appg && adduser -S appu -G appg && apk add --no-cache curl

USER appu

COPY --from=builder /app/pingApp . 

ENTRYPOINT [ "/pingApp" ]