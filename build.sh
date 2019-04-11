#!/usr/bin/env bash


platforms=("linux/amd64" "linux/386" "darwin/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name='sf-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi  
    output_name+='.tar.gz'

    env GOOS=$GOOS GOARCH=$GOARCH go build -o sf $package
    tar -czf $output_name sf
    mkdir -p artifacts
    mv $output_name artifacts/$output_name
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done