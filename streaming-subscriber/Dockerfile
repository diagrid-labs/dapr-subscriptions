FROM cgr.dev/chainguard/go AS builder

WORKDIR /work

COPY go.mod /work/
COPY go.sum /work/
COPY cmd /work/cmd

RUN CGO_ENABLED=0 go build -o streamingsub ./cmd/main.go

FROM cgr.dev/chainguard/static
COPY --from=builder /work/streamingsub /streamingsub
CMD ["/streamingsub"]
