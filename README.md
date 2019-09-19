# verge-adapter

verge-adapter继承了bitcoin-adapter，主要修改了如下内容：

- 重写了Symbol = "XVG"。
- 重写了addressDecoder，实现了XVG地址编码。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建XVG.ini文件，编辑如下内容：

```ini

# RPC Server Type，0: CoreWallet RPC; 1: Explorer API
rpcServerType = 0
# node api url
serverAPI = "http://"
# RPC Authentication Username
rpcUser = "walletUser"
# RPC Authentication Password
rpcPassword = "walletPassword2019"
# Is network test?
isTestNet = false
# support segWit
supportSegWit = false
# minimum transaction fees
minFees = "0.2"

```

## 官方资料

### github

https://github.com/vergecurrency/VERGE

### 浏览器

https://verge-blockchain.info/