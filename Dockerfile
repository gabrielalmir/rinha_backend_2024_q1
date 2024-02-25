FROM golang:1.21.5 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/api

FROM scratch
COPY --from=build /app/api /app/api
ENTRYPOINT [ "/app/api" ]
