FROM golang:1.18.3
WORKDIR /app

COPY .implementation .
COPY .server .
EXPOSE 50051
CMD ["go", "run", "server/main.go"]