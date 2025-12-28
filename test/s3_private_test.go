package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

// NOTE: Valid AWS credentials are required to run this test.
// Run with: go test -v -timeout 30m s3_private_test.go

func TestS3PrivateModule(t *testing.T) {
	t.Parallel()

	nameSuffix := random.UniqueId()
	// S3 buckets must be globally unique and lowercase
	bucketName := fmt.Sprintf("terratest-bucket-%s", strings.ToLower(nameSuffix))

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../s3-private",

		Vars: map[string]interface{}{
			"bucket_name": bucketName,
			"versioning":  true,
		},

		Lock: false,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndPlan(t, terraformOptions)
}
