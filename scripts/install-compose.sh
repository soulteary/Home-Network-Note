#!/bin/bash

# author @soulteary
# date 2022/02/10

echo "Downloading and installing docker-compose..."
BASEURL="https://github.com/docker/compose/releases/download"
FILENAME=$(echo "docker-compose-`uname -s`-`uname -m`" | tr [:upper:] [:lower:])
VERSION=2.2.3

if [ -f "$FILENAME" ]; then
    echo "$FILENAME exists.";
else
    URL_COMPOSE="${BASEURL}/v${VERSION}/${FILENAME}"
    echo "[Download] ${URL_COMPOSE}"
    curl -L "${URL_COMPOSE}" -o ${FILENAME}
fi

curl -L "${BASEURL}/v${VERSION}/${FILENAME}.sha256" -o ${FILENAME}.sha256
CHECK=$(shasum -c ${FILENAME}.sha256)

if [ $? != 0 ]; then
    echo "${FILENAME} checksum is not valid";
    exit 1;
fi

mv ${FILENAME} /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
rm ${FILENAME}.sha256

docker-compose -v
echo "done";
