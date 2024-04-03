package provider

import (
	"fmt"
	"github.com/Regis24GmbH/terraform-provider-s3bucketnotification/internal/pkg/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var s3client = client.S3Client{}

func TestAccResourceS3BucketNotification(t *testing.T) {
	os.Setenv("TF_ACC", "1")
	os.Setenv("TF_CLI_CONFIG_FILE", "/dev/null")
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"awsr24": New(),
		},
		CheckDestroy: testAccCheckS3BucketNotificationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccS3BucketNotificationConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckS3BucketNotificationExists("awsr24_s3bucketnotification.example"),
					// Add more checks as needed
				),
			},
			// Add more steps as needed to test updates
		},
	})
}

func testAccPreCheck(t *testing.T) {
	// Implement this function to check that the necessary API credentials and setup has been done
}

func testAccCheckS3BucketNotificationDestroy(s *terraform.State) error {
	// Implement this function to check that the S3 bucket notification has been destroyed
	return nil
}

func testAccCheckS3BucketNotificationExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Implement this function to check that the S3 bucket notification exists
		return nil
	}
}

var testAccS3BucketNotificationConfig = fmt.Sprintf(`provider "awsr24" {
  access_key = "%s"
  secret_key = "%s"
  region     = "eu-west-1"
  alias = "awsr24"
}

resource "awsr24_s3bucketnotification" "example" {
  bucket = "my-bucket"

  queue_configurations {
    queue_arn = "arn:aws:sqs:us-west-2:123456789012:myqueue"
    events = ["s3:ObjectCreated:*"]
    filter {
      filter_rules {
        name = "prefix"
        value = "images/"
      }
      filter_rules {
        name = "suffix"
        value = ".jpg"
      }
    }
    id = "MyQueueConfig"
  }
}
`, "accesskey", "secretkey")
