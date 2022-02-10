FROM golang:1.16-alpine

WORKDIR /code

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

CMD [ "/code/bin/air" ]
