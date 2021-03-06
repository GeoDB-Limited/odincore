FROM golang:1.13.5-buster as builder

WORKDIR /go-owasm
COPY go-owasm/. /go-owasm

WORKDIR /chain
COPY chain/ /chain

RUN wget https://github.com/WebAssembly/wabt/releases/download/1.0.17/wabt-1.0.17-ubuntu.tar.gz
RUN tar -zxf wabt-1.0.17-ubuntu.tar.gz
RUN cp wabt-1.0.17/bin/wat2wasm /usr/local/bin

RUN make install && make faucet

FROM debian:latest

RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

RUN apt-get install wget

COPY --from=builder /go-owasm/api/libgo_owasm.so /lib/libgo_owasm.so
COPY --from=builder /go/bin/bandd /usr/local/bin/bandd
COPY --from=builder /go/bin/bandcli /usr/local/bin/bandcli
COPY --from=builder /go/bin/yoda /usr/local/bin/yoda
COPY --from=builder /go/bin/faucet /usr/local/bin/faucet

COPY chain/docker-config/validator1/ validator1/
COPY chain/docker-config/validator2/ validator2/
COPY chain/docker-config/validator3/ validator3/
COPY chain/docker-config/validator4/ validator4/

# generated genesis
COPY chain/docker-config/genesis.json genesis.json

CMD ["bandd", "--help"]
