#!/bin/sh
# author soulteary

command_exists() {
    command -v "$@" >/dev/null 2>&1
}

use_curl() {
    command_exists curl || return 1
}

use_wget() {
    command_exists wget || return 1
}

declare -a files
files[0]=https://ftp.arin.net/pub/stats/arin/delegated-arin-extended-latest
files[1]=https://ftp.arin.net/pub/stats/arin/delegated-arin-extended-latest.md5
files[2]=https://ftp.ripe.net/ripe/stats/delegated-ripencc-extended-latest
files[3]=https://ftp.ripe.net/ripe/stats/delegated-ripencc-extended-latest.md5
files[4]=https://ftp.apnic.net/stats/apnic/delegated-apnic-extended-latest
files[5]=https://ftp.apnic.net/stats/apnic/delegated-apnic-extended-latest.md5
files[6]=https://ftp.lacnic.net/pub/stats/lacnic/delegated-lacnic-extended-latest
files[7]=https://ftp.lacnic.net/pub/stats/lacnic/delegated-lacnic-extended-latest.md5
files[8]=https://ftp.afrinic.net/pub/stats/afrinic/delegated-afrinic-extended-latest
files[9]=https://ftp.afrinic.net/pub/stats/afrinic/delegated-afrinic-extended-latest.md5

download() {
    if use_curl; then
        for file in ${files[*]}
        do
            curl -O "${file}"
        done
    else
        if use_wget; then
            for file in ${files[*]}
            do
                wget -c "${file}"
            done
        else
            echo "You need install `wget` or `curl` first."
            return 1
        fi
    fi
}

download
