# PKSF-Private-Testnet-Product-Tracebility

### Container
```
docker ps --format "table {{.Names}}\t{{.Image}}\t{{.Ports}}"
```

### Orderer
```
docker logs -f orderer.example.com
```

### Peers
```
docker logs -f peer0.org1.example.com
docker logs -f peer0.org2.example.com
```

### Fabric CAs
```
docker logs -f ca_org1
docker logs -f ca_org2
docker logs -f ca_orderer
```

### Chaincode runtime containers
```
docker logs -f dev-peer0.org1.example.com-fisheries_1.0-f7e50d2...
docker logs -f dev-peer0.org2.example.com-livestock_1.0-d92e...
```

### Your REST API app
```
docker logs -f musing_buck
```
