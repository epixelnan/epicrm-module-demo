FROM golang:1.20
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
RUN CGO_ENABLED=0 GOOS=linux go build -o apiserver

FROM alpine
COPY --from=0 /app/src/apiserver /app/
WORKDIR /app/
CMD ["./apiserver"]
