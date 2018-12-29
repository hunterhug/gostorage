#!/bin/bash
docker build -t seaweed:latest .
docker tag seaweed:latest hunterhug/seaweed:latest