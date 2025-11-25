# https://playwright-community.github.io/playwright-go/

# Stage 1: Modules caching
FROM golang:1.25.3 AS modules

COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Stage 2: Build
FROM golang:1.25.3 AS builder

COPY --from=modules /go/pkg /go/pkg
COPY . /workdir

WORKDIR /workdir

# Install playwright cli with right version for later use
RUN PWGO_VER=$(grep -oE "playwright-go v\S+" /workdir/go.mod | sed 's/playwright-go //g') \
    && go install github.com/playwright-community/playwright-go/cmd/playwright@${PWGO_VER}

# Stage 3: Final
FROM ubuntu:noble

COPY --from=builder /go/bin/playwright /

COPY --from=builder /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

COPY . /workdir

WORKDIR /workdir

RUN apt-get update
RUN apt-get install -y ca-certificates tzdata
    # Install dependencies and all browsers (or specify one)
RUN /playwright install --with-deps
RUN rm -rf /var/lib/apt/lists/*

CMD ["go", "test", "-v", "./test/e2e/..."]