mkdir -p src/github.com/ethereum
pushd src/github.com/ethereum
git clone https://github.com/ethereum/go-ethereum.git
cd go-ethereum
git reset --hard 1e248f3a6e14f3bfc9ebe1b315c4194300220f68
make
popd
export GOPATH=$PWD
go build
