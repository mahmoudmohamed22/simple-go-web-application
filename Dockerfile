FROM golang:1.18-alpine
RUN mkdir /app
COPY ./go-app/. /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]



