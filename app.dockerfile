FROM golang:1.21.1-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o reddit cmd/web/*.go && ./reddit 

EXPOSE 8080

CMD [ "/reddit" ]