FROM golang:1.15 as builder

COPY . /src

WORKDIR /src
RUN CGO_ENABLED=0 go build -ldflags="-w -s"

FROM scratch

COPY --from=builder /src/universal-redirect-https /

EXPOSE 80
ENTRYPOINT [ "/universal-redirect-https" ]

LABEL org.opencontainers.image.source https://github.com/riksby/universal-redirect-https
