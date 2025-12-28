package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

// NOTE: Valid AWS credentials are required to run this test.
// Run with: go test -v -timeout 30m security_group_test.go

func TestSecurityGroupModule(t *testing.T) {
	t.Parallel()

	nameSuffix := random.UniqueId()
	sgName := fmt.Sprintf("test-sg-%s", nameSuffix)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../security-group",

		Vars: map[string]interface{}{
			"name":                sgName,
			"vpc_id":              "vpc-mock-id",
			"ingress_cidr_blocks": []string{"0.0.0.0/0"},
			"ingress_rules":       []string{"http-80-tcp"},
		},

		Lock: false,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndPlan(t, terraformOptions)
}
