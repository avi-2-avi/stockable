EC2_PUBLIC_IP=$(curl -s --retry 3 --connect-timeout 5 http://169.254.169.254/latest/meta-data/public-ipv4)

if [[ -z "$EC2_PUBLIC_IP" ]]; then
    echo "Error: Failed to retrieve public IP. Defaulting to localhost."
    EC2_PUBLIC_IP="127.0.0.1"
fi

echo "Public IP: $EC2_PUBLIC_IP"
echo "VITE_API_URL=http://$EC2_PUBLIC_IP:8085" > /home/ec2-user/stockable/frontend/.env.local
chown ec2-user:ec2-user /home/ec2-user/stockable/frontend/.env.local

NGINX_CONF="/home/ec2-user/stockable/frontend/nginx.conf"
sed -i "s|http://stockable-backend:8085/|http://$EC2_PUBLIC_IP:8085/|g" "$NGINX_CONF"

DOCKER_COMPOSE_FILE="/home/ec2-user/stockable/docker-compose.yml"
sed -i "s|http://stockable-backend:8085|http://$EC2_PUBLIC_IP:8085|g" "$DOCKER_COMPOSE_FILE"

BACKEND_ENV="/home/ec2-user/stockable/backend/.env"
if grep -q "^ALLOWED_ORIGIN=" "$BACKEND_ENV"; then
    sed -i "s|^ALLOWED_ORIGIN=.*|ALLOWED_ORIGIN=http://$EC2_PUBLIC_IP|" "$BACKEND_ENV"
else
    echo "ALLOWED_ORIGIN=http://$EC2_PUBLIC_IP" >> "$BACKEND_ENV"
fi

SECURITY_GROUP_ID=$(aws ec2 describe-instances \
  --query "Reservations[*].Instances[*].SecurityGroups[0].GroupId" \
  --output text 2>/dev/null)

if [[ -z "$SECURITY_GROUP_ID" ]]; then
    echo "Error: Failed to retrieve Security Group ID."
else
    aws ec2 authorize-security-group-ingress \
        --group-id "$SECURITY_GROUP_ID" \
        --protocol tcp \
        --port 8085 \
        --cidr "${EC2_PUBLIC_IP}/32" 2>/dev/null

    if [[ $? -ne 0 ]]; then
        echo "Warning: Failed to update security group. Check permissions."
    else
        echo "Security group updated successfully."
    fi
fi

echo "Running Docker Compose..."
cd /home/ec2-user/stockable
docker-compose up -d