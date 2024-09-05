FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o upfast-tf2-web .

EXPOSE 8080

CMD ["/app/upfast-tf2-web", "-prefork=false", "-dev=false", "-port=:8080"]

