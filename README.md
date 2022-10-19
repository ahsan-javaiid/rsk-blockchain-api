# rsk-blockchain-api

- Restapi to interact with rskj node. This api uses goroutines and channels to handle parallel requests and scalable enough to handle multiple requests at same time
- I have plan to add websockets support. Idea is to keep checking the new blocks, transactions and logs in a loop and emit the events

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
- net_peerList
 ```
GET localhost:8080/testnet/net_peerList

// Result
{
  "id":1,
  "jsonrpc": "2.0",
  "result": [
       "8d19eb7e9c21484ea831557d63a6a80146ca2469af2b4044125dc8b7598b3b6dca5ad14a2c693a70f572f060be435d5d2822464a126817bffc5623746fe8d8c3 | 190.57.233.165/190.57.233.165:60480",
       ...
  ]
}
 ```

 - eth_chainId
 ```
GET localhost:8080/testnet/eth_chainId

// Result
{
    "jsonrpc": "2.0",
    "id": 67,
    "result": "0x1f"  // in hex format
}
 ```

### More Coming soon ....