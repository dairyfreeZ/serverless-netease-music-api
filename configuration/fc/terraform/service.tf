locals {
  log_project_name      = "nm-api-log-proj"
  log_store_name        = "nm-api-log-store"
  service_name          = "nm-api-service"
  service_ram_role_name = "nm-api-service-ram-role"
}

resource "alicloud_fc_service" "default" {
  name       = local.service_name
  role       = alicloud_ram_role.default.arn
  depends_on = [alicloud_ram_role_policy_attachment.default]
}

resource "alicloud_ram_role" "default" {
  name     = local.service_ram_role_name
  document = <<EOF
        {
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Effect": "Allow",
              "Principal": {
                "Service": [
                  "fc.aliyuncs.com"
                ]
              }
            }
          ],
          "Version": "1"
        }

EOF
  force    = true
}

resource "alicloud_ram_role_policy_attachment" "default" {
  role_name   = alicloud_ram_role.default.name
  policy_name = "AliyunLogFullAccess"
  policy_type = "System"
}
