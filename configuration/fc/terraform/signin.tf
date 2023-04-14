locals {
  fc_name     = "signin-fc"
  filename    = "signin_fc.zip"
  binary_name = "signin"
}

resource "alicloud_fc_function" "signin_fc" {
  service     = alicloud_fc_service.default.name
  name        = local.fc_name
  filename    = local.filename
  handler     = "signin"
  memory_size = "128"
  description = "test"

  runtime = "go1"
}

data "archive_file" "fc" {
  type        = "zip"
  source_file = "../entry/release/${local.binary_name}"
  output_path = local.filename
}
