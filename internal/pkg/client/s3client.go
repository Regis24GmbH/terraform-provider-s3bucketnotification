package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	svc *s3.S3
}

// NewClient creates a new S3Client with session
func NewClient(sess *session.Session) *S3Client {
	svc := s3.New(sess)

	return &S3Client{
		svc: svc,
	}
}

func (c *S3Client) CreateBucketNotification(bucket string, notificationConfiguration *s3.PutBucketNotificationConfigurationInput) error {
	_, err := c.svc.PutBucketNotificationConfiguration(notificationConfiguration)
	return err
}

func (c *S3Client) GetBucketNotification(bucket string) (*s3.NotificationConfiguration, error) {
	result, err := c.svc.GetBucketNotificationConfiguration(&s3.GetBucketNotificationConfigurationRequest{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *S3Client) UpdateBucketNotification(bucket string, notificationConfiguration *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
	return c.svc.PutBucketNotificationConfiguration(notificationConfiguration)
}

func (c *S3Client) DeleteBucketNotification(bucket string) error {
	_, err := c.svc.PutBucketNotificationConfiguration(&s3.PutBucketNotificationConfigurationInput{
		Bucket:                    aws.String(bucket),
		NotificationConfiguration: &s3.NotificationConfiguration{},
	})
	return err
}
