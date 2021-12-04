go build -o miningPoolCli main.go
mkdir miningPoolCli-unix-x86–64

cp miningPoolCli miningPoolCli-unix-x86–64/
cp LICENSE miningPoolCli-unix-x86–64/
cp README.md miningPoolCli-unix-x86–64/

tar -zcvf miningPoolCli-unix-x86–64.tar.gz miningPoolCli-unix-x86–64/