FROM golang:1.18.3-bullseye

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /coinstream

EXPOSE 8000

CMD [ "/coinstream" ]