module "security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "5.3.1"

  name        = var.name
  description = "Security group for ${var.name}"
  vpc_id      = var.vpc_id

  ingress_cidr_blocks = var.ingress_cidr_blocks
  ingress_rules       = var.ingress_rules
  egress_rules        = ["all-all"]
}
