# Ethereum Development with Go

รายการ Live! Code ครั้งที่ 36 เรื่อง Ethereum Development with Go ผมจะโค้ดเรื่องการพัฒนา Go Programming เรียกใช้งาน Smart Contracts บน Ethereum Blockchain 
- Ethereum Blockchain
- Smart Contracts with Solidity
- Smart Contracts Unit Test
- ERC-20
- Development with Go

ดูย้อนหลังได้ที่ https://fb.com/CodeBangkok/videos

## Setting up the Client (macOS)
1. ติดตั้ง Solidity Compiler
```
brew update
brew tap ethereum/ethereum
brew install solidity
```
2. ติดตั้ง abigen
```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

## Smart Contract Complilation & Create Go contract file

```
solc --abi --bin bondcoin.sol -o build
abigen --bin=./build/bondcoin.bin --abi=./build/bondcoin.abi --pkg=bondcoin --out=bondcoin.go
```

อ้างอิง https://goethereumbook.org
