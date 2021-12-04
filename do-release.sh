go build -o miningPoolCli main.go
mkdir miningPoolCli_unix_x86_64

cp miningPoolCli miningPoolCli_unix_x86_64/
cp LICENSE miningPoolCli_unix_x86_64/
cp README.md miningPoolCli_unix_x86_64/

tar -zcvf miningPoolCli_unix_x86_64.tar.gz miningPoolCli_unix_x86_64/