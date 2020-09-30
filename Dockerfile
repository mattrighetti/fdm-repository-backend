FROM golang:1.15.2-alpine3.12

LABEL org.opencontainers.image.source https://github.com/MattRighetti/fdm-repository-backend

WORKDIR $GOPATH/src/fdm-repository-backend

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["fdm-repository-backend"]