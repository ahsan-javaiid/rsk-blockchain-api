# rsk-blockchain-api

Rest api to interact with rskj node 

### Starting the server
```
go run main.go
```

### Supported RPC calls via http rest interface
#### URL format
```
GET http://localhost:8080/{mainnet|testnet}/{rpc-method-name}
```

#### Example URL's

- web3_clientVersion
```
GET http://localhost:8080/testnet/web3_clientVersion

// Result
{
    "jsonrpc": "2.0",
    "id": 67,
    "result": "RskJ/4.0.0/Linux/Java1.8/HOP-08d8acd"
}
``` 
- net_version
 ```
GET http://localhost:8080/testnet/net_version

// Result
{
    "jsonrpc": "2.0",
    "id": 67,
    "result": "30"
}
 ```
 - net_peerCount
 ```
GET localhost:8080/testnet/net_peerCount

// Result
{
  "id":74,
  "jsonrpc": "2.0",
  "result": "0x2" // 2
}
 ```

### More Coming soon ....