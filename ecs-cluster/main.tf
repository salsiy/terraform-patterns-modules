module "ecs_service" {
  source  = "terraform-aws-modules/ecs/aws//modules/service"
  version = "6.11.0"

  name        = var.service_name
  cluster_arn = var.cluster_arn

  cpu    = var.cpu
  memory = var.memory

  enable_execute_command = true

  container_definitions = var.container_definitions

  subnet_ids = var.subnet_ids

  launch_type = "FARGATE"
  network_mode = "awsvpc"

}
