# Terraform Patterns Modules

This repository contains reusable Terraform modules for the "Production-Grade Terraform Patterns" series.

## Modules

- **ecs-cluster**: Deploys an ECS Cluster and defines Services (Dry/Consolidated pattern).
- **alb**: Wrapper for AWS ALB.
- **s3-private**: Wrapper for Secure S3 Buckets.
- **security-group**: Wrapper for Security Groups.

## Releases

This repository uses **Release Please**.
- Commits with `feat:` or `fix:` triggers a Release PR.
- Merging the Release PR creates a new Tag (e.g., `v1.1.0`) and updates `CHANGELOG.md`.

## Usage
Reference these modules in your `terraform-patterns-live` repository using the `git` source URL with a tag.
```hcl
source = "git::https://github.com/salsiy/terraform-patterns-modules.git//ecs-cluster?ref=v1.0.0"
```