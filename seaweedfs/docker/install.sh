#!/bin/sh

mkdir -p /storage/app/seaweed/filer
chmod 777 /storage/app/seaweed/filer
mkdir -p /storage/app/seaweed/master
mkdir -p /storage/app/seaweed/volume1
mkdir -p /storage/app/seaweed/volume2

docker-compose up -d
