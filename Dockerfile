# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.18-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

# Get dependencies - will also be cached if we won't change go.mod/go.sum
COPY go.mod /go-xpayments/
COPY go.sum /go-xpayments/
RUN cd /go-xpayments && go mod download

ADD . /go-xpayments
RUN cd /go-xpayments && go run build/ci.go install ./cmd/gpay

# Pull Gpay into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-xpayments/build/bin/gpay /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["gpay"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"
