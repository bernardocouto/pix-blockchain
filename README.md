# pix-blockchain
PIX Blockchain

## Docker

Inicialização da rede:
```shell
cd blockchain
docker-compose -f docker-compose.yaml up
```

## Chaincode

Construção e inicialização dos chaincode.

Acesso ao terminal do container do chaincode:
```shell
docker exec -it chaincode sh
```

Compilação dos chaincodes:
```shell
cd qr_code_estatico
go build -o qr_code_estatico
cd ../qr_code_dinamico
go build -o qr_code_dinamico
```

Executar os chaincode:
```shell
CORE_CHAINCODE_LOGLEVEL=debug CORE_CHAINCODE_ID_NAME=qr_code_estatico:0 CORE_PEER_TLS_ENABLED=false ./qr_code_estatico -peer.address peer:7052
CORE_CHAINCODE_LOGLEVEL=debug CORE_CHAINCODE_ID_NAME=qr_code_dinamico:0 CORE_PEER_TLS_ENABLED=false ./qr_code_dinamico -peer.address peer:7052
```

## Utilizando os chaincodes

Acesso ao terminal do container do CLI:
```shell
docker exec -it cli bash
```
