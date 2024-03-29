FROM golang:1.18 AS development

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /app

EXPOSE 8081

CMD ["app"]