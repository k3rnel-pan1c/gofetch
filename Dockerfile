FROM golang:1.26

WORKDIR /app

COPY src/ .

RUN go build -o gofetch .

CMD ["bash"]
