terraform {
  required_version = ">= 0.13"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.12.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

resource "aws_iam_policy" "dynamodb_policy" {
  name        = var.dynamodb_policy_name
  path        = "/"
  description = "Policy to grant permissions to access the DynamoDB table"

  policy = data.dynamodb_policy_document
}

resource "aws_dynamodb_table" "database" {
  name           = var.dynamodb_table_name
  billing_mode   = "PROVISIONED"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "Url"

  attribute {
    name = "Url"
    type = "S"
  }

  ttl {
    attribute_name = "TimeToLive"
    enabled        = true
  }

  tags = {
    Name        = var.dynamodb_table_name
    Environment = var.environment
  }
}
