FROM golang:1.12-alpine AS base
WORKDIR /src
RUN apk add --no-cache git

FROM base AS lint-and-test
RUN go install golang.org/x/lint/golint@latest
COPY ./ ./
RUN golint ./...
RUN go test -mod vendor -v 2>&1 ./statuspage/...

FROM base AS build
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod=vendor -ldflags="-s -w" -a -o build/terraform-provider-statuspage-darwin-amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -mod=vendor -ldflags="-s -w" -a -o build/terraform-provider-statuspage-linux-386
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags="-s -w" -a -o build/terraform-provider-statuspage-linux-amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -mod=vendor -ldflags="-s -w" -a -o build/terraform-provider-statuspage-linux-arm
RUN for binary in build/*; do sha256sum -b $binary > $binary.sha256; done