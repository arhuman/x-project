# build stage
FROM golang:1.22 as builder

# setup the working directory
WORKDIR /api

# install dependencies
COPY go.*  /api/
RUN go mod download

# add source code
COPY cmd /api/cmd
COPY doc /api/doc
COPY static /api/static
COPY internal /api/internal
COPY models /api/models

# build the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/api/api.go
# RUN go build -o server ./cmd/api/server.go

# use a minimal alpine image
FROM alpine
# In prod use distroless
# FROM gcr.io/distroless/base

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# set working directory
WORKDIR /app

# copy the binary from builder
COPY --from=builder /api/server .
COPY .env /app/.env
RUN source /app/.env

EXPOSE 8080

# run the binary
CMD ["./server"]
