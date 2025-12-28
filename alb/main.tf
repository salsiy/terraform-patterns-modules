module "alb" {
  source  = "terraform-aws-modules/alb/aws"
  version = "10.4.0"

  name    = var.name
  vpc_id  = var.vpc_id
  subnets = var.subnets

  enable_deletion_protection = false

  security_groups = var.security_groups

  listeners = {
    http-https-redirect = {
      port     = 80
      protocol = "HTTP"
      redirect = {
        port        = "443"
        protocol    = "HTTPS"
        status_code = "HTTP_301"
      }
    }
  }

}
