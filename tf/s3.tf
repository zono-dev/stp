// bucket for stp
resource "aws_s3_bucket" "s3_for_image" {
  bucket = "${var.s3.bucket_name_prefix}-${random_string.lower.result}"
  acl    = "private"
  tags = {
    Name = var.tag.tag_name
  }
}

// log bucket for accessing s3_for_image
resource "aws_s3_bucket" "log_bucket" {
  bucket = "log-${var.s3.bucket_name_prefix}-${random_string.lower.result}"
  acl    = "log-delivery-write"
}

// notification setting for stp lambda
resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = aws_s3_bucket.s3_for_image.id
  lambda_function {
    lambda_function_arn = aws_lambda_function.stp_lambda.arn
    events              = ["s3:ObjectCreated:Put"]
    filter_prefix       = var.s3.origin_folder_path
  }
  depends_on = [aws_lambda_permission.allow_bucket]
}

// public access setting for s3_for_image and log_bucket
resource "aws_s3_bucket_public_access_block" "s3_for_image" {
  bucket                  = aws_s3_bucket.s3_for_image.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_public_access_block" "log_bucket" {
  bucket                  = aws_s3_bucket.log_bucket.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

// making folder for stp working
resource "aws_s3_bucket_object" "s3_for_image_org" {
  bucket = aws_s3_bucket.s3_for_image.id
  acl    = "private"
  key    = var.s3.origin_folder_path
  source = "resource/0bytes.txt"
}

resource "aws_s3_bucket_object" "s3_for_image_resize" {
  bucket = aws_s3_bucket.s3_for_image.id
  acl    = "private"
  key    = var.s3.resized_folder_path
  source = "resource/0bytes.txt"
}

// setup the bucket policy
resource "aws_s3_bucket_policy" "bucket_policy_for_image_bucket" {
  bucket = aws_s3_bucket.s3_for_image.id
  policy = data.aws_iam_policy_document.s3_policy_for_stp.json
}
