# Build the Go program
FROM golang:1.23 AS builder
RUN apt-get update && apt-get install -y wget xz-utils
RUN cd /usr/local/bin && wget -O- https://github.com/upx/upx/releases/download/v5.0.1/upx-5.0.1-amd64_linux.tar.xz | tar -Jxvf- --strip-components=1 upx-5.0.1-amd64_linux/upx

WORKDIR /app
COPY . .
RUN go mod init main
RUN CGO_ENABLED=0 go build -o main -ldflags="-w -s" .
RUN upx --best --lzma main

# Create a minimal runtime image
FROM scratch
EXPOSE 1323
COPY --from=builder /app/main /main
CMD ["/main"]