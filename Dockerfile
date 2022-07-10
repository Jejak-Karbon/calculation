FROM golang:alpine

WORKDIR /var/www/html/apps/calculation
COPY . /var/www/html/apps/calculation

RUN go build -o main .

CMD ["/var/www/html/apps/calculation/main"]