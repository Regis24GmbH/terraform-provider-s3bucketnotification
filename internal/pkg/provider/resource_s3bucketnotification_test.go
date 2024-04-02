package provider_test

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"testing"

	"github.com/Regis24GmbH/terraform-provider-s3bucketnotification/internal/pkg/provider"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockS3Client struct {
	mock.Mock
}

func (m *mockS3Client) GetBucketNotification(bucket string) (*s3.NotificationConfiguration, error) {
	args := m.Called(bucket)
	return args.Get(0).(*s3.NotificationConfiguration), args.Error(1)
}

func (m *mockS3Client) UpdateBucketNotification(bucket string, notificationConfiguration *s3.PutBucketNotificationConfigurationInput) error {
	args := m.Called(bucket, notificationConfiguration)
	return args.Error(0)
}

func TestResourceS3BucketNotificationCreate(t *testing.T) {
	s3Client := new(mockS3Client)
	s3Client.On("GetBucketNotification", "test-bucket").Return(&s3.NotificationConfiguration{}, nil)
	s3Client.On("UpdateBucketNotification", "test-bucket", mock.Anything).Return(nil)

	resourceData := schema.TestResourceDataRaw(t, provider.ResourceS3BucketNotificationSchema().Schema, map[string]interface{}{
		"bucket": "test-bucket",
	})

	diags := provider.ResourceS3BucketNotificationSchema().CreateContext(context.Background(), resourceData, s3Client)

	assert.Equal(t, diag.Diagnostics(nil), diags)
	s3Client.AssertExpectations(t)
}

func TestResourceS3BucketNotificationCreate_ErrorGettingBucketNotification(t *testing.T) {
	s3Client := new(mockS3Client)
	s3Client.On("GetBucketNotification", "test-bucket").Return(nil, awserr.New(s3.ErrCodeNoSuchBucket, "Bucket does not exist", nil))

	resourceData := schema.TestResourceDataRaw(t, provider.ResourceS3BucketNotificationSchema().Schema, map[string]interface{}{
		"bucket": "test-bucket",
	})

	diags := provider.ResourceS3BucketNotificationSchema().CreateContext(context.Background(), resourceData, s3Client)

	assert.NotEqual(t, diag.Diagnostics(nil), diags)
	s3Client.AssertExpectations(t)
}

func TestResourceS3BucketNotificationCreate_ErrorUpdatingBucketNotification(t *testing.T) {
	s3Client := new(mockS3Client)
	s3Client.On("GetBucketNotification", "test-bucket").Return(&s3.NotificationConfiguration{}, nil)
	s3Client.On("UpdateBucketNotification", "test-bucket", mock.Anything).Return(awserr.New(s3.ErrCodeBucketAlreadyExists, "Bucket already exists", nil))

	resourceData := schema.TestResourceDataRaw(t, provider.ResourceS3BucketNotificationSchema().Schema, map[string]interface{}{
		"bucket": "test-bucket",
	})

	diags := provider.ResourceS3BucketNotificationSchema().CreateContext(context.Background(), resourceData, s3Client)

	assert.NotEqual(t, diag.Diagnostics(nil), diags)
	s3Client.AssertExpectations(t)
}
