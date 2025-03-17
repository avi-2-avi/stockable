output "ssh_command" {
  value = <<EOT
ssh -i "${var.key_name}.pem" ec2-user@${aws_instance.stockable_ec2.public_dns}
EOT
  description = "Use this command to SSH into your EC2 instance"
}
