provider "aws" {
  region = var.region
}

resource "aws_lambda_function" "singin_lambda" {
  filename      = "signin_lambda.zip"
  function_name = "signin_lambda"
  handler       = "signin"
  role          = aws_iam_role.signin_lambda_exec_role.arn
  description   = "test3"

  runtime = "go1.x"
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "signin_lambda_exec_role" {
  name               = "signin_lambda_exec_role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "archive_file" "lambda" {
  type        = "zip"
  source_file = "../entry/release/signin"
  output_path = "signin_lambda.zip"
}
