data "dynamodb_policy_document" "dynamodb_policy_data" {
  statement {
    actions = [
      "dynamodb:PutItem",
      "dynamodb:GetItem",
    ]

    resources = [
      "arn:aws:dynamodb:${var.aws_region}:&{aws:userid}:table/${var.dynamodb_table_name}"
    ]
  }
}
