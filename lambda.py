import boto3
import logging
from botocore.exceptions import ClientError

# Initialize clients for S3 and CloudWatch
s3 = boto3.client('s3')
cloudwatch = boto3.client('cloudwatch')

# Define the S3 bucket and file to check
bucket_name = 'my-unique-bucket-name'  # Replace with your bucket name
file_name = 'your-file-name.txt'  # Replace with the file name you're checking

def handler(event, context):
    metric_value = 1  # Default to 1 (file not found)
    
    try:
        # Check if the file exists in the S3 bucket
        s3.head_object(Bucket=bucket_name, Key=file_name)
        metric_value = 0  # File found, set metric value to 0
    except ClientError as e:
        # If the error is 'NotFound', continue with the default metric value
        if e.response['Error']['Code'] != 'NoSuchKey':
            logging.error(f"Error checking file: {e}")
            raise e
    
    # Publish the metric to CloudWatch using the previously created custom metric
    cloudwatch.put_metric_data(
        Namespace='Custom/FileCheck',  # The namespace of the previously created custom metric
        MetricData=[
            {
                'MetricName': 'FileExistence',  # The metric name from the previous custom metric
                'Dimensions': [
                    {'Name': 'BucketName', 'Value': bucket_name},
                    {'Name': 'FileName', 'Value': file_name}
                ],
                'Value': metric_value,
                'Unit': 'None'
            }
        ]
    )
    logging.info(f"Metric published: {metric_value}")
