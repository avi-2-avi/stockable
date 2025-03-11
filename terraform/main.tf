resource "aws_security_group" "stockable_sg" {
  name        = "stockable-security-group"
  description = "Allow inbound traffic for Stockable"

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # HTTP
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "stockable_ec2" {
  ami           = var.ec2_ami
  instance_type = var.instance_type
  key_name      = var.key_name
  security_groups = [aws_security_group.stockable_sg.name]

  user_data = file("${path.module}/user_data.sh")

  tags = {
    Name = "StockableApp"
  }
}
