package provider

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/Regis24GmbH/terraform-provider-s3bucketnotification/internal/pkg/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"log"
)

func resourceS3BucketNotification() *schema.Resource {
	return ResourceS3BucketNotificationSchema()
}

func resourceS3BucketNotificationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Implement your resource creation logic here
	log.Printf("[INFO] Create S3BucketNotification: %s", d.Get("bucket").(string))
	sess := meta.(*session.Session)

	s3Client := client.NewClient(sess)
	bucketNotificationConfiguration, err := s3Client.GetBucketNotification(d.Get("bucket").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	eventBridgeConfiguration := s3.EventBridgeConfiguration{}
	lambdaFunctionConfigurations := make([]*s3.LambdaFunctionConfiguration, 0)
	queueConfigurations := make([]*s3.QueueConfiguration, 0)
	topicConfigurations := make([]*s3.TopicConfiguration, 0)

	if bucketNotificationConfiguration != nil {
		eventBridgeConfiguration = *bucketNotificationConfiguration.EventBridgeConfiguration
		lambdaFunctionConfigurations = bucketNotificationConfiguration.LambdaFunctionConfigurations
		queueConfigurations = bucketNotificationConfiguration.QueueConfigurations
		topicConfigurations = bucketNotificationConfiguration.TopicConfigurations
	}

	queueConfigurations = appendQueueConfigurationsFromSchema(d, queueConfigurations)
	lambdaFunctionConfigurations = appendLambdaConfigurationsFromSchema(d, lambdaFunctionConfigurations)
	topicConfigurations = appendTopicConfigurationsFromSchema(d, topicConfigurations)

	output, err := s3Client.UpdateBucketNotification(d.Get("bucket").(string), &s3.PutBucketNotificationConfigurationInput{
		Bucket: aws.String(d.Get("bucket").(string)),
		NotificationConfiguration: &s3.NotificationConfiguration{
			EventBridgeConfiguration:     &eventBridgeConfiguration,
			LambdaFunctionConfigurations: lambdaFunctionConfigurations,
			QueueConfigurations:          queueConfigurations,
			TopicConfigurations:          topicConfigurations,
		},
	})

	if err != nil {
		return diag.FromErr(err)
	}

	// create a sha hash from a string
	sha := sha256.New()
	sha.Write([]byte(output.GoString()))
	shaHash := hex.EncodeToString(sha.Sum(nil))
	d.SetId(shaHash)

	return nil
}

func appendQueueConfigurationsFromSchema(d *schema.ResourceData, queueConfigurations []*s3.QueueConfiguration) []*s3.QueueConfiguration {
	for _, queueConfiguration := range d.Get("queue_configurations").([]interface{}) {
		queueConfigurations = append(queueConfigurations, &s3.QueueConfiguration{
			Events:   aws.StringSlice(queueConfiguration.(map[string]interface{})["events"].([]string)),
			Filter:   createFilter(queueConfiguration.(map[string]interface{})["filter"].([]map[string]string)),
			Id:       aws.String(queueConfiguration.(map[string]interface{})["id"].(string)),
			QueueArn: aws.String(queueConfiguration.(map[string]interface{})["queue_arn"].(string)),
		})
	}
	return queueConfigurations
}

func appendLambdaConfigurationsFromSchema(d *schema.ResourceData, lambdaFunctionConfigurations []*s3.LambdaFunctionConfiguration) []*s3.LambdaFunctionConfiguration {
	for _, lambdaFunctionConfiguration := range d.Get("lambda_function_configurations").([]interface{}) {
		lambdaFunctionConfigurations = append(lambdaFunctionConfigurations, &s3.LambdaFunctionConfiguration{
			Events:            aws.StringSlice(lambdaFunctionConfiguration.(map[string]interface{})["events"].([]string)),
			Filter:            createFilter(lambdaFunctionConfiguration.(map[string]interface{})["filter"].([]map[string]string)),
			Id:                aws.String(lambdaFunctionConfiguration.(map[string]interface{})["id"].(string)),
			LambdaFunctionArn: aws.String(lambdaFunctionConfiguration.(map[string]interface{})["lambda_function_arn"].(string)),
		})
	}
	return lambdaFunctionConfigurations
}

func appendTopicConfigurationsFromSchema(d *schema.ResourceData, topicConfigurations []*s3.TopicConfiguration) []*s3.TopicConfiguration {
	for _, topicConfiguration := range d.Get("topic_configurations").([]interface{}) {
		topicConfigurations = append(topicConfigurations, &s3.TopicConfiguration{
			Events:   aws.StringSlice(topicConfiguration.(map[string]interface{})["events"].([]string)),
			Filter:   createFilter(topicConfiguration.(map[string]interface{})["filter"].([]map[string]string)),
			Id:       aws.String(topicConfiguration.(map[string]interface{})["id"].(string)),
			TopicArn: aws.String(topicConfiguration.(map[string]interface{})["topic_arn"].(string)),
		})
	}
	return topicConfigurations
}

// function to create the filter for all configurations from the schema
func createFilter(filterRules []map[string]string) *s3.NotificationConfigurationFilter {
	filter := &s3.NotificationConfigurationFilter{}
	if len(filterRules) > 0 {
		filter.Key = &s3.KeyFilter{
			FilterRules: make([]*s3.FilterRule, 0),
		}
		for _, rule := range filterRules {
			filter.Key.FilterRules = append(filter.Key.FilterRules, &s3.FilterRule{
				Name:  aws.String(rule["name"]),
				Value: aws.String(rule["value"]),
			})
		}
	}
	return filter
}

func resourceS3BucketNotificationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Implementieren Sie hier die Logik zum Lesen Ihrer Ressource
	log.Printf("[INFO] Read S3BucketNotification: %s", d.Id())
	return nil
}

func resourceS3BucketNotificationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Implementieren Sie hier die Logik zum Aktualisieren Ihrer Ressource
	log.Printf("[INFO] Update S3BucketNotification: %s", d.Id())
	return nil
}

func resourceS3BucketNotificationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Implementieren Sie hier die Logik zum Löschen Ihrer Ressource
	log.Printf("[INFO] Delete S3BucketNotification: %s", d.Id())
	return nil
}
