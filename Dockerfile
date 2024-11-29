FROM golang

COPY . .

RUN go build -o riot .

CMD ["./riot"]
