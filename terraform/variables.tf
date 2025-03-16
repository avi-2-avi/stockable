variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t3.micro"
}

variable "key_name" {
  description = "AWS key pair name for SSH access"
  type        = string

}

variable "ssh_admin_ip" {
  description = "Your public IP address (IPv4) for SSH access. Leave empty to disable SSH."
  type        = string
}


variable "security_group_name" {
  description = "Security group name"
  default     = "stockable-sg"
  type        = string
}

variable "ec2_ami" {
  description = "Amazon Machine Image (AMI) ID for the EC2 instance"
  type        = string
  default     = "ami-00d4cdb3bda21c6ed"
}

variable "private_key_path" {
  description = "Path to the private key file used for SSH"
  type        = string
}