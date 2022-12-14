FROM golang:1.19 AS grpc-health-probe-builder

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.6 && \
    wget -qO/bin/grpc_health_probe \
    https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

FROM golang:1.19 AS builder

COPY . /src
WORKDIR /src

RUN make build

FROM debian:stable-slim

RUN apt-get update &&  \
    apt-get install -y --no-install-recommends  \
    ca-certificates netbase &&  \
    rm -rf /var/lib/apt/lists/ && \
    apt-get autoremove -y && apt-get autoclean -y

COPY --from=grpc-health-probe-builder /bin/grpc_health_probe /usr/local/bin/
COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 9000 8000

VOLUME /data/conf

CMD ["./bin", "-conf", "/data/conf"]
