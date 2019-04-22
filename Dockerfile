FROM circleci/golang:1.12 AS base
WORKDIR /src

FROM base
USER root
RUN go get -u golang.org/x/lint/golint
COPY ./ ./
RUN golint ./...
RUN go test -mod vendor -v 2>&1 ./statuspage/...

FROM base AS build
USER root
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod=vendor -ldflags="-s -w" -a -o build/terraform-provider-statuspage-darwin-amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -mod=vendor -ldflags="-s -w" -a -o build/terraform-provider-statuspage-linux-386
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags="-s -w" -a -o build/terraform-provider-statuspage-linux-amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -mod=vendor -ldflags="-s -w" -a -o build/terraform-provider-statuspage-linux-arm
RUN for binary in build/*; do sha256sum -b $binary > $binary.sha256; done