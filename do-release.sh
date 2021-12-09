PARSE_VER=`awk '/BuildVersion/{print $NF}' config/version.go`
BUILD_VERSION=${PARSE_VER:1:-1}

FOLDER="miningPoolCli-${BUILD_VERSION}"
TAR_NAME="miningPoolCli-${BUILD_VERSION}-linux.tar.gz"

printf "Creating release v${BUILD_VERSION}\n\n"

go build -o miningPoolCli main.go

mkdir $FOLDER
touch "${FOLDER}/VERSION_${BUILD_VERSION}_x86_x64"

cp miningPoolCli LICENSE README.md $FOLDER
cp hiveos_configs/* $FOLDER
sed -i -e "s/CUSTOM_VERSION=/CUSTOM_VERSION=${BUILD_VERSION}/g" $FOLDER/h-manifest.conf

tar -zcvf "${TAR_NAME}" $FOLDER
rm -rf $FOLDER