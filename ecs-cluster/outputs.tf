output "service_id" {
  description = "ARN of the ECS Service"
  value       = module.ecs_service.id
}

output "task_definition_arn" {
  description = "ARN of the Task Definition"
  value       = module.ecs_service.task_definition_arn
}
