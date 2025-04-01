FROM golang:1.21 AS builder
WORKDIR /app
RUN git clone https://github.com/cosmos/cosmos-sdk.git && \
    cd cosmos-sdk && \
    git checkout v0.47.5 && \
    make build
RUN mv /app/cosmos-sdk/build/simd /usr/bin/simd

FROM ubuntu:20.04
COPY --from=builder /usr/bin/simd /usr/bin/simd
RUN apt-get update && apt-get install -y curl
EXPOSE 26656 26657 1317 9090
ENTRYPOINT ["simd"]
CMD ["start"]