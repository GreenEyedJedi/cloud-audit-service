resource "random_id" "bucket_id" {
  byte_length = 4
}

resource "aws_s3_bucket" "audit_data" {
  bucket        = "alec-cloud-audit-data-${random_id.bucket_id.hex}"
  force_destroy = true

  tags = {
    Project = "CloudAuditService"
    Owner   = "Alec"
  }
}