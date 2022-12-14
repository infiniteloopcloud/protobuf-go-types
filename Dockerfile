# Go builder
FROM golang:1.19-alpine3.17 as go_builder
RUN apk add --no-cache gcc libc-dev 
WORKDIR /out

ENV CGO_ENABLED=0
ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/infiniteloopcloud/protoc-gen-go-types@latest
COPY . .
RUN go build -o /out/protoc-gen-go-types .

# Final
FROM alpine:3.17

COPY --from=go_builder /out/protoc-gen-go-types /usr/bin/protoc-gen-go-types
RUN apk add --no-cache libstdc++ protobuf

ENTRYPOINT ["/usr/bin/protoc"]