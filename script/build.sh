#!/bin/bash
aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 805199394057.dkr.ecr.us-east-2.amazonaws.com
docker build -t mbora/goserver .
docker tag mbora/goserver:latest 805199394057.dkr.ecr.us-east-2.amazonaws.com/mbora/goserver:latest
docker push 805199394057.dkr.ecr.us-east-2.amazonaws.com/mbora/goserver:latest