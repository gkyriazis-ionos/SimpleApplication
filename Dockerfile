FROM golang:1.18.3
WORKDIR /app

COPY .implementation .
COPY .server .
#COPY .client .

EXPOSE 50051
CMD ["go", "run", "server/main.go"]
#CMD ["go", "run", "client/main.go", "-name=\"GOPATH\""]