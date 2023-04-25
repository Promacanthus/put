FROM docker-reg.devops.xiaohongshu.com/shequ/golang:1.19.2 as builder

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o put cmd/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM docker-reg.devops.xiaohongshu.com/shequ/alpine-3.13:tz

WORKDIR /workspace

COPY --from=builder /workspace/put .

ENTRYPOINT ["/workspace/put"]