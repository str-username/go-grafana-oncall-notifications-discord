FROM golang:1.21.3-alpine3.18 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -a -installsuffix cgo -o notify cmd/main.go


FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /app/

COPY --from=build /app/notify .

COPY etc/users.json .

CMD ["./notify"]