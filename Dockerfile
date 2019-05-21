FROM golang:1.12 as foundation

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM foundation as builder
ENV CGO_ENABLED=0
COPY . .
RUN go test -v ./... \
    && go build -a -tags netgo -ldflags '-s -w "-extldflags=-static"' -o bin/vergen

FROM gcr.io/distroless/base as runtime

COPY --from=builder /build/bin/vergen /bin/vergen

ENTRYPOINT ["/bin/vergen"]
