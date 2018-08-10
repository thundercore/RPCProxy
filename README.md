# RPCProxy

## About

At Thunder, we found ourselves wanting a tool to help us monitor and log the details of our RPC calls. We forked [@ymedialabs' ReverseProxy](https://github.com/ymedialabs/ReverseProxy) and made it pick out the different RPC calls being made.

## Prerequisites

 - Go Version 1.6+
     - Install via https://golang.org/doc/install

## Getting Started

After cloning this repository, run the build script with `sh build.sh`. This will clone the [`go-ethereum`](https://github.com/ethereum/go-ethereum/) repository into your directory (under `src`), adjust your `GOPATH`, and build the binary for RPCProxy.

## Running RPCProxy

By default, RPCProxy will be running on `http://localhost:9999` and it will be forwarding the requests to http://testnet-rpc.thundercore.com:8545. You can specify the forwarding "port" and "url"  by using the `-port` and `-url` command line options.

Running the binary:
```
./RPCProxy
```

Setting new parameters to run the binary with:
```
./RPCProxy -port=9090 -url=http://other_rpc:8545
```

Outputting the contents to a log file:
```
./RPCProxy 2>log.txt
```

## Output

When you run RPCProxy, it will print the port where the proxy is running and where it is redirecting the request to.

When requests start being made, it will print the RPC call and parameters used, the HTTP headers, the raw transaction data, and the response to `stderr`. It should look something like this:

```
eth_getBalance [0x9a78d67096ba0c7c1bcdc0a8742649bc399119c0 latest]
2018/08/06 17:23:57 ----------------New Request-------------------------
 {"id":5289471075661898,"jsonrpc":"2.0","params":["0x9a78d67096ba0c7c1bcdc0a8742649bc399119c0","latest"],"method":"eth_getBalance"}  
    Response Time 107301342 
    Response Body------
 HTTP/1.1 200 OK
Content-Length: 76
Access-Control-Allow-Origin: chrome-extension://nkbihfbeogaeaoehlefnkodbefgpgknn
Content-Type: application/json
Date: Mon, 06 Aug 2018 21:23:57 GMT
Vary: Origin

{"jsonrpc":"2.0","id":5289471075661898,"result":"0x49fb0a92b073f3c1a2f6cc"}
```
