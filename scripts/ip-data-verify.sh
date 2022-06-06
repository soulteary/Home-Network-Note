#!/bin/sh
# author soulteary

command_exists() {
    command -v "$@" >/dev/null 2>&1
}

use_md5() {
    command_exists md5 || return 1
}

use_md5sum() {
    command_exists md5sum || return 1
}


files=$(ls data/*latest)

download() {
    if use_md5; then
        for file in ${files[*]}
        do
            echo $file
            digst=$(md5 -r $file | awk '{print $1}')
            checksum=$(cat "$file".md5 | awk '{print $4}')
            if [ "$digst" != "$checksum" ]; then
                checksum=$(cat "$file".md5 | awk '{print $1}')
                if [ "$digst" != "$checksum" ]; then
                    echo "$file data validation failed, please try downloading again."
                fi
            fi
        done
    else
        if use_md5sum; then
            for file in ${files[*]}
            do
                echo $file
                digst=$(md5sum $file | awk '{print $1}')
                checksum=$(cat "$file".md5 | awk '{print $4}')
                if [ "$digst" != "$checksum" ]; then
                    checksum=$(cat "$file".md5 | awk '{print $1}')
                    if [ "$digst" != "$checksum" ]; then
                        echo "$file data validation failed, please try downloading again."
                    fi
                fi
            done
        else
            echo "You need install `md5` or `md5sum` first."
            return 1
        fi
    fi
}

download
