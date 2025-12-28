variable "name" {
  description = "Name of security group"
  type        = string
}

variable "vpc_id" {
  description = "ID of the VPC"
  type        = string
}

variable "ingress_cidr_blocks" {
  description = "List of IPv4 CIDR ranges to use on all ingress rules"
  type        = list(string)
  default     = []
}

variable "ingress_rules" {
  description = "List of ingress rules to create"
  type        = list(string)
  default     = ["http-80-tcp", "https-443-tcp"]
}
