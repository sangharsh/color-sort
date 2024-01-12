# Inspiration: https://github.com/grpc-ecosystem/grpc-cloud-run-example/blob/master/golang/README.md
FROM golang:alpine3.19 as build

WORKDIR /go-server
ENV HOST 0.0.0.0
COPY proto/game.proto ./proto/
COPY server/go.mod .
COPY server/go.sum .
COPY server/ ./

# Installs protoc and required plugin
RUN apk add protoc && \
    go get -u google.golang.org/protobuf/cmd/protoc-gen-go && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go && \
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc


ARG SRC_DIR=./proto
ARG GO_DST_DIR=./gen
RUN mkdir -p $GO_DST_DIR && \
    protoc \
    --proto_path==${SRC_DIR} \
    --go_out=${GO_DST_DIR} \
    --go_opt=module=github.com/sangharsh/color-sort \
    --go-grpc_out=${GO_DST_DIR} \
    --go-grpc_opt=module=github.com/sangharsh/color-sort \
    ${SRC_DIR}/game.proto

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo \
    -o /go/bin/server \
    github.com/sangharsh/color-sort


FROM scratch

COPY --from=build /go/bin/server /server
ENV HOST 0.0.0.0
ENTRYPOINT ["/server"]
