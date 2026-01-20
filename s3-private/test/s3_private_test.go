package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3Private(t *testing.T) {
	t.Parallel()

	// Generate a random bucket name to avoid collisions
	bucketName := fmt.Sprintf("test-bucket-%s", strings.ToLower(random.UniqueId()))

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located (one directory up)
		TerraformDir: "../",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"bucket_name": bucketName,
		},

		// Set environment variables
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": "us-east-1",
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	outputArn := terraform.Output(t, terraformOptions, "bucket_arn")

	// Verify we're getting back a valid ARN
	assert.Contains(t, outputArn, "arn:aws:s3")
	assert.Contains(t, outputArn, bucketName)
}
