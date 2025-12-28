output "lb_id" {
  description = "The ID of the load balancer"
  value       = module.alb.id
}

output "lb_arn" {
  description = "The ARN of the load balancer"
  value       = module.alb.arn
}

output "lb_dns_name" {
  description = "The DNS name of the load balancer"
  value       = module.alb.dns_name
}
