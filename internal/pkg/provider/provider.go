package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_ACCESS_KEY_ID", nil),
				Description: "The access key for API operations. You can retrieve this from the 'Security & Identity' section of the AWS console.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_SECRET_ACCESS_KEY", nil),
				Description: "The secret key for API operations. You can retrieve this from the 'Security & Identity' section of the AWS console.",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The region to send requests to. This should be a valid AWS region such as 'us-east-1'.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"awsr24_s3bucketnotification": resourceS3BucketNotification(),
		},
	}
}
