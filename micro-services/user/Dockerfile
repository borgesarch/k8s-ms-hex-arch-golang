FROM golang:1.16.4
WORKDIR /app
COPY . .
RUN go mod download
COPY /.env /app/
RUN go build -o main ./bin
EXPOSE 8842
CMD ["./main"]