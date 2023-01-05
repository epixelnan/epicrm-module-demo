FROM golang:1.17
COPY src/go.mod /app/src/go.mod
COPY src/go.sum /app/src/go.sum
WORKDIR /app/src/
RUN go mod download -x

WORKDIR /
COPY src /app/src

WORKDIR /app/src/

RUN go get
RUN go test ./...
# The generated binary fails to execute if cgo is enabled
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine
COPY --from=0 /app/src/epicmod_hello_helloapi /app/
WORKDIR /app/
CMD ["./epicmod_hello_helloapi"]
