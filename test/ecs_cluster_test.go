package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

// NOTE: Valid AWS credentials are required to run this test.
// Run with: go test -v -timeout 30m ecs_cluster_test.go

func TestEcsClusterModule(t *testing.T) {
	t.Parallel()

	nameSuffix := random.UniqueId()
	serviceName := fmt.Sprintf("test-service-%s", nameSuffix)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../ecs-cluster",

		Vars: map[string]interface{}{
			"service_name": serviceName,
			"cluster_arn":  "arn:aws:ecs:us-east-1:123456789012:cluster/main-cluster",
			"subnet_ids":   []string{"subnet-mock-1", "subnet-mock-2"},
			"container_definitions": map[string]interface{}{
				"main": map[string]interface{}{
					"image":  "nginx:latest",
					"cpu":    256,
					"memory": 512,
					"portMappings": []map[string]interface{}{
						{
							"containerPort": 80,
							"hostPort":      80,
							"protocol":      "tcp",
						},
					},
				},
			},
		},

		Lock: false,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndPlan(t, terraformOptions)
}
