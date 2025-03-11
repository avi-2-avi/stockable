#!/bin/bash
yum update -y

yum install -y docker
service docker start
usermod -aG docker ec2-user
systemctl enable docker

curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

git clone ${docker_compose_repo} /home/ec2-user/stockable
cd /home/ec2-user/stockable

docker-compose up -d
