FROM --platform=linux/amd64 golang:1.22

WORKDIR /app

# Copy dependencies
COPY .env go.mod go.sum ./

RUN go mod download

# Copy src
COPY ./challenges/ ./challenges/
COPY ./dotenv/ ./dotenv/

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /mugen-api

EXPOSE 8080

CMD ["/mugen-api"]
