FROM golang:alpine AS build

WORKDIR /app
COPY . .
RUN go build ./cmd/main.go

ENTRYPOINT [ "/app/main" ][]

# to create image
# docker build -t string-svc .
# to run image
# docker run -d -p 8080:8080 string-svc