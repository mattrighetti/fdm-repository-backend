############################
# STEP 1 build executable  #
############################
FROM golang:1.15.2-alpine3.12 AS build-phase
LABEL org.opencontainers.image.source https://github.com/mattrighetti/fdm-repository-backend
WORKDIR $GOPATH/src/fdm-repository-backend
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/fdm-backend .

############################
# STEP 2  create image     #
############################
FROM scratch
COPY --from=build-phase /go/bin/fdm-backend /go/bin/fdm-backend
EXPOSE 8080
ENTRYPOINT ["/go/bin/fdm-backend"]