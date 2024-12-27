FROM --platform=linux/amd64 golang:1.23

WORKDIR /app

COPY ./start.sh /start.sh

RUN chmod +x /start.sh
RUN go install github.com/cespare/reflex@latest

ENTRYPOINT [ "/start.sh" ]

CMD []