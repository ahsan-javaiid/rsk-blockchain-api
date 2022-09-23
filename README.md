# rsk-blockchain-api

Rest api to interact with rskj node 

### Starting the server
```
go run main.go
```

### Supported RPC calls via http rest interface
#### URL format
http://localhost:8080/{mainnet|testnet}/{rpc-method-name}

#### Example URL's
- http://localhost:8080/mainnet/web3_clientVersion 
- http://localhost:8080/mainnet/net_version
