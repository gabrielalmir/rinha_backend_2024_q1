## Build
FROM golang:1.22.0 AS build
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/api

## Deploy
FROM scratch

COPY --from=build /app/api /app/api

ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT [ "/app/api" ]
