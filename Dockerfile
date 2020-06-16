# builder image
FROM golang:1.14-buster as builder

WORKDIR /src
COPY . .
RUN go build -o /bin/walletd -v -ldflags "-w -s"

# final image
FROM debian:buster

ENV XDG_CONFIG_HOME /etc/xdg

RUN apt-get update && apt-get install -y \
  ca-certificates \
  dumb-init

COPY --from=builder /bin/walletd /bin/walletd

USER 65534
ENTRYPOINT ["dumb-init", "--", "/bin/walletd"]
