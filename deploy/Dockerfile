#docker build -t backend-golang .
#docker run -dt -p 8080:8080 --name backend-go backend-golang

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /iBP-goapp

EXPOSE 8080

CMD [ "/CIBP" ]