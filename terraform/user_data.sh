#!/bin/bash
set -e

yum update -y
yum install -y git docker aws-cli
service docker start
usermod -aG docker ec2-user
systemctl enable docker

curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

mkdir -p /home/ec2-user/stockable
chown ec2-user:ec2-user /home/ec2-user/stockable

sudo -u ec2-user git clone https://github.com/avi-2-avi/stockable.git /home/ec2-user/stockable
chown -R ec2-user:ec2-user /home/ec2-user/stockable

cd /home/ec2-user/stockable
sudo -u ec2-user git checkout main

aws s3 cp s3://stockable-env-files/backend.env /home/ec2-user/stockable/backend/.env