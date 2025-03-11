variable "aws_region" {
  description = "AWS region where resources will be created"
  type        = string
  default     = "us-east-1"
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t3.micro"
}

variable "key_name" {
  description = "Name of the AWS key pair for SSH access"
  type        = string
}

variable "ec2_ami" {
  description = "Amazon Machine Image (AMI) ID for the EC2 instance"
  type        = string
  default     = "ami-0905a3c97561e0b69" 
}

variable "docker_compose_repo" {
  description = "GitHub repo URL containing the Docker Compose setup"
  type        = string
}
