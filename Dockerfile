FROM golang:1.23

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

COPY .env .

RUN go build -o ecommerce-order

RUN chmod +x ecommerce-order

EXPOSE 9002

CMD [ "./ecommerce-order" ]