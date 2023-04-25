FROM golang:1.20 as builder

COPY . .

RUN go mod download
# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o put cmd/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static-debian11

WORKDIR /workspace

COPY --from=builder /workspace/put .

ENTRYPOINT ["/workspace/put"]