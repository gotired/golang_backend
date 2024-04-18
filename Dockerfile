FROM golang:1.22.2-alpine AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM scratch

COPY --from=build /app/main /app/main
COPY --from=build /app/.env /.env

CMD ["/app/main"]
