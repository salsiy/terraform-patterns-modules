variable "service_name" {
  description = "Name of the ECS service"
  type        = string
}

variable "cluster_arn" {
  description = "ARN of the ECS Cluster"
  type        = string
}

variable "cpu" {
  description = "CPU units for the task"
  type        = number
  default     = 256
}

variable "memory" {
  description = "Memory (in MiB) for the task"
  type        = number
  default     = 512
}

variable "container_definitions" {
  description = "Map of container definitions"
  type        = any
}

variable "subnet_ids" {
  description = "List of subnet IDs"
  type        = list(string)
}
