# make sure we exit if an error occurs
set -e

if [ "$1" == "" ]
then
    echo Please specify package version, e.g.  0.1.0
    read PACKAGE_VERSION
else
    PACKAGE_VERSION=$1
fi

rm -rf ../docs && rm -rf ../hwmux_client
curl --compressed https://hwmux.silabs.net/schema/download -o hwmux.yaml
docker compose run --rm openapi-generator-cli generate -i /local/tools/hwmux.yaml -g go --additional-properties=projectName=hwmux-client-golang,packageVersion=$PACKAGE_VERSION,packageName=hwmux -o /local/
