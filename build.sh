#!/usr/bin/env bash

go build -o file-server
docker build -t file-server:0.1 .
