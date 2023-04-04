#Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
#RUN go mod download
#RUN go build -o main main.go

RUN go get .
RUN go build -o main main.go

#Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main" ]