BUILD_VERSION="1.0.11"
FOLDER="miningPoolCli-${BUILD_VERSION}"
TAR_NAME="miningPoolCli-${BUILD_VERSION}-linux.tar.gz"

go build -o miningPoolCli main.go

mkdir $FOLDER
touch "${FOLDER}/VERSION_${BUILD_VERSION}_x86_x64"

cp miningPoolCli LICENSE README.md $FOLDER
cp hiveos_configs/* $FOLDER

tar -zcvf "${TAR_NAME}" $FOLDER
rm -rf $FOLDER