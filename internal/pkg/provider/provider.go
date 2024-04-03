package provider

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
				DefaultFunc: schema.EnvDefaultFunc("AWS_DEFAULT_REGION", "eu-west-1"),
				Description: "The region to send requests to. This should be a valid AWS region such as 'us-east-1'.",
			},
			"role_arn": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_DEFAULT_REGION", "eu-west-1"),
				Description: "The region to send requests to. This should be a valid AWS region such as 'us-east-1'.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"awsr24_s3bucketnotification": resourceS3BucketNotification(),
		},
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	accessKey := d.Get("access_key").(string)
	secretKey := d.Get("secret_key").(string)
	region := d.Get("region").(string)
	roleArn := d.Get("role_arn").(string) // The ARN of the role to assume

	config := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	// Assume the role
	creds := stscreds.NewCredentials(sess, roleArn)

	// Create a new session with the assumed role
	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: creds,
	})

	if err != nil {
		return nil, diag.FromErr(err)
	}

	return sess, nil
}
