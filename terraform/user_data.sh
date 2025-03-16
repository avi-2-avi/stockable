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
sudo -u ec2-user git checkout feat/deploy-v2

aws s3 cp s3://stockable-env-files/backend.env /home/ec2-user/stockable/backend/.env

EC2_PUBLIC_IP=$(curl -s ifconfig.me)
echo "VITE_API_URL=http://$EC2_PUBLIC_IP:8085" > /home/ec2-user/stockable/frontend/.env.local
chown ec2-user:ec2-user /home/ec2-user/stockable/frontend/.env.local

SECURITY_GROUP_ID=$(aws ec2 describe-instances \
  --query "Reservations[*].Instances[*].SecurityGroups[0].GroupId" \
  --output text)

aws ec2 authorize-security-group-ingress \
    --group-id $SECURITY_GROUP_ID \
    --protocol tcp \
    --port 8085 \
    --cidr ${EC2_PUBLIC_IP}/32

# docker-compose up -d