# https://playwright-community.github.io/playwright-go/

# Stage 1: Modules caching
FROM golang:1.25.2 AS modules

COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Stage 2: Build
FROM golang:1.25.2 AS builder

COPY --from=modules /go/pkg /go/pkg
COPY . /workdir

WORKDIR /workdir

# Install playwright cli with right version for later use
RUN PWGO_VER=$(grep -oE "playwright-go v\S+" /workdir/go.mod | sed 's/playwright-go //g') \
    && go install github.com/playwright-community/playwright-go/cmd/playwright@${PWGO_VER}

# Build your app
# RUN GOOS=linux GOARCH=amd64 go build -o /bin/myapp

# Stage 3: Final
FROM ubuntu:noble

COPY --from=builder /go/bin/playwright /

COPY --from=builder /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

RUN apt-get update && apt-get install -y ca-certificates tzdata \
    # Install dependencies and all browsers (or specify one)
    && /playwright install --with-deps \
    && rm -rf /var/lib/apt/lists/*
# CMD ["/myapp"]
CMD ["go", "test", "-v", "./test/e2e/..."]

# # Step 1: Modules caching
# FROM golang:1.25.2-alpine3.21 AS modules

# COPY go.mod go.sum /modules/

# WORKDIR /modules

# RUN go run github.com/playwright-community/playwright-go/cmd/playwright@v0.5200.1 install --with-deps

# RUN go mod download

# # Step 2: Tests
# FROM golang:1.25.2-alpine3.21

# COPY --from=modules /go/pkg /go/pkg
# COPY . /app

# WORKDIR /app

# RUN go env -w CGO_ENABLED=0 \
#     && go env -w GOOS=linux  \
#     && go env -w GOARCH=amd64

# CMD ["go", "test", "-v", "./test/e2e/..."]