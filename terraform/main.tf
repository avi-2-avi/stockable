resource "aws_s3_bucket" "stockable_env_bucket" {
  bucket = "stockable-env-files"
}

resource "aws_s3_object" "backend_env" {
  bucket = aws_s3_bucket.stockable_env_bucket.id
  key    = "backend.env"
  source = "${path.module}/../backend/.env"
  acl    = "private"
}

resource "aws_s3_object" "frontend_env_local" {
  bucket = aws_s3_bucket.stockable_env_bucket.id
  key    = "frontend.env.local"
  source = "${path.module}/../frontend/.env.local"
  acl    = "private"
}

resource "aws_iam_role" "stockable_ec2_role" {
  name = "stockable-ec2-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "ec2.amazonaws.com"
      }
    }]
  })
}

resource "aws_iam_policy" "s3_access_policy" {
  name        = "StockableS3Access"
  description = "Allows EC2 to read .env files from S3"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect   = "Allow"
      Action   = ["s3:GetObject"]
      Resource = "arn:aws:s3:::stockable-env-files/*"
    }]
  })
}

resource "aws_iam_role_policy_attachment" "attach_s3_policy" {
  role       = aws_iam_role.stockable_ec2_role.name
  policy_arn = aws_iam_policy.s3_access_policy.arn
}

resource "aws_iam_instance_profile" "stockable_ec2_profile" {
  name = "stockable-ec2-profile"
  role = aws_iam_role.stockable_ec2_role.name
}

resource "aws_security_group" "stockable_sg" {
  name        = var.security_group_name
  description = "Allow inbound traffic for Stockable"

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # HTTP
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # HTTPS
  }

  ingress {
    from_port   = 8085
    to_port     = 8085
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Stockable API
  }

  dynamic "ingress" {
    for_each = length(regexall("^\\d+\\.\\d+\\.\\d+\\.\\d+$", var.ssh_admin_ip)) > 0 ? [1] : []
    content {
      from_port   = 22
      to_port     = 22
      protocol    = "tcp"
      cidr_blocks = ["${var.ssh_admin_ip}/32"]
    }
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "stockable_ec2" {
  ami             = var.ec2_ami
  instance_type   = var.instance_type
  key_name        = var.key_name
  security_groups = [aws_security_group.stockable_sg.name]

  iam_instance_profile = aws_iam_instance_profile.stockable_ec2_profile.name

  user_data = file("${path.module}/user_data.sh")

  tags = {
    Name = "StockableApp"
  }
}
