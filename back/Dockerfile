FROM golang:1.17

ADD . /backend

WORKDIR /backend
RUN go mod init back
RUN go mod tidy
RUN go build -o back .
RUN chmod +x /backend

ENTRYPOINT ["/backend/back"]