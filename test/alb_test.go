package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

// NOTE: Valid AWS credentials are required to run this test.
// Run with: go test -v -timeout 30m alb_test.go

func TestAlbModule(t *testing.T) {
	t.Parallel()

	// Append a random suffix to ensure unique names
	nameSuffix := random.UniqueId()
	albName := fmt.Sprintf("test-alb-%s", nameSuffix)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../alb",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":            albName,
			"vpc_id":          "vpc-mock-id", // Needs real VPC ID for apply, mock for plan
			"subnets":         []string{"subnet-mock-1", "subnet-mock-2"},
			"security_groups": []string{"sg-mock-1"},
		},

		// Disable locking for local testing
		Lock: false,
	})

	// At the end of the test, run `terraform destroy`
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`
	// Note: Verify will likely verify plan-only in environment without creds
	// For real testing, use InitAndApply
	terraform.InitAndPlan(t, terraformOptions)
	// terraform.InitAndApply(t, terraformOptions)

	// Validate output if apply were successful
	// output := terraform.Output(t, terraformOptions, "alb_arn")
	// assert.Contains(t, output, "arn:aws:elasticloadbalancing")
}
