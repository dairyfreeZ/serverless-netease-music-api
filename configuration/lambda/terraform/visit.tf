resource "aws_lambda_function" "visit_lambda" {
  filename         = "visit_lambda.zip"
  source_code_hash = data.archive_file.visit_lambda_archive.output_base64sha256
  function_name    = "visit_lambda"
  handler          = "visit"
  role             = aws_iam_role.visit_lambda_exec_role.arn

  runtime = "go1.x"
}

resource "aws_iam_role" "visit_lambda_exec_role" {
  name               = "visit_lambda_exec_role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role_policy_attachment" "visit_lambda_exec_role_permissions_attach" {
  policy_arn = aws_iam_policy.lambda_exec_role_permissions.arn
  role       = aws_iam_role.visit_lambda_exec_role.name
}

data "archive_file" "visit_lambda_archive" {
  type        = "zip"
  source_file = "../entry/release/visit"
  output_path = "visit_lambda.zip"
}
