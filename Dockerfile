FROM golang

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY * ./

RUN go build -o /docker-goapi

CMD ["/docker-goapi"]