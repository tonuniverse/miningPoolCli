BUILD_VERSION="v1.0.3-beta"
FOLDER="miningPoolCli-${BUILD_VERSION}-unix-x86-64"
TAR_NAME="${FOLDER}.tar.gz"

go build -o miningPoolCli main.go

mkdir $FOLDER
cp miningPoolCli LICENSE README.md $FOLDER
tar -zcvf $TAR_NAME $FOLDER
rm -rf $FOLDER