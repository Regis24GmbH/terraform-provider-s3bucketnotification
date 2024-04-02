## Usage

```
provider "awsr24" {
  access_key = "my-access-key"
  secret_key = "my-secret-key"
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

  lambda_function_configurations {
    lambda_function_arn = "arn:aws:lambda:us-west-2:123456789012:function:MyLambdaFunction"
    events = ["s3:ObjectCreated:*", "s3:ObjectRemoved:*"]
    filter {
      filter_rules {
        name = "prefix"
        value = "pdfs/"
      }
    }
    id = "MyLambdaFunctionConfig"
  }

  topic_configurations {
    topic_arn = "arn:aws:sns:us-west-2:123456789012:MyTopic"
    events = ["s3:ObjectCreated:*"]
    filter {
      filter_rules {
        name = "prefix"
        value = "logs/"
      }
      filter_rules {
        name = "suffix"
        value = ".log"
      }
    }
    id = "MyTopicConfig"
  }
}
```

```
provider "awsr24" {
  access_key = "my-access-key"
  secret_key = "my-secret-key"
}

resource "awsr24_s3bucketnotification" "example" {
  bucket = "my-bucket"

  queue_configurations {
    queue_arn = "arn:aws:sqs:us-west-2:123456789012:myqueue1"
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
    id = "MyQueueConfig1"
  }

  queue_configurations {
    queue_arn = "arn:aws:sqs:us-west-2:123456789012:myqueue2"
    events = ["s3:ObjectRemoved:*"]
    filter {
      filter_rules {
        name = "prefix"
        value = "logs/"
      }
      filter_rules {
        name = "suffix"
        value = ".log"
      }
    }
    id = "MyQueueConfig2"
  }
}
```
multiple configurations can be added to the same bucket notification resource. The above example creates a bucket notification with a queue, lambda function, and SNS topic configuration.
