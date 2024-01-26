FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o reddit cmd/web/*.go && ./reddit 

EXPOSE 8080

CMD [ "/reddit" ]