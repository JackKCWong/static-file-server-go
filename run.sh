#!/usr/bin/env bash

docker run --rm -it -v /mnt/d/workspace/tianchi/tail-based-sampling/data:/mnt/data -p 8888:8888 -e "SERVER_PORT=8888" --name "file-server" file-server:0.1
