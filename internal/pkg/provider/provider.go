package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"s3bucketnotification_s3bucketnotification": resourceS3BucketNotification(),
		},
	}
}
