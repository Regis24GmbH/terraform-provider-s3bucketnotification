package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func ResourceS3BucketNotificationSchema() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceS3BucketNotificationCreate,
		ReadContext:   resourceS3BucketNotificationRead,
		UpdateContext: resourceS3BucketNotificationUpdate,
		DeleteContext: resourceS3BucketNotificationDelete,

		Schema: map[string]*schema.Schema{
			"bucket": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"queue_configurations": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     queueConfigurationSchema(),
			},
			"lambda_function_configurations": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     lambdaFunctionConfigurationSchema(),
			},
			"topic_configurations": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     topicConfigurationSchema(),
			},
		},
	}
}

func queueConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"queue_arn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"events": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"filter": filterSchema(),
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func lambdaFunctionConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"lambda_function_arn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"events": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"filter": filterSchema(),
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func topicConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"topic_arn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"events": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"filter": filterSchema(),
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func filterSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"filter_rules": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": &schema.Schema{
								Type:     schema.TypeString,
								Optional: true,
							},
							"value": &schema.Schema{
								Type:     schema.TypeString,
								Optional: true,
							},
						},
					},
				},
			},
		},
	}
}
