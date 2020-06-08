locals {
  lambda_function_name = "${var.lambda.function_name_prefix}-${random_string.lower.result}"
  file_path = "../bin/function.zip"
}


resource "aws_lambda_function" "stp_lambda" {
  filename      = local.file_path
  function_name = local.lambda_function_name
  handler       = var.lambda.handler
  source_code_hash = filebase64sha256(local.file_path)
  timeout       = var.lambda.timeout
  memory_size   = var.lambda.memory_size
  role          = aws_iam_role.role_for_stp.arn
  depends_on    = [aws_cloudwatch_log_group.stp_log_group]
  runtime       = "go1.x"
  environment {
    variables = {
      VLOG    = var.lambda.environment.VLOG
      DTABLE  = aws_dynamodb_table.table_for_image.name
      PREFIX  = var.lambda.environment.PREFIX
      RESIZEX = var.lambda.environment.RESIZEX
      RESIZEY = var.lambda.environment.RESIZEY
    }
  }
}

resource "aws_cloudwatch_log_group" "stp_log_group" {
  name              = "/aws/lambda/${local.lambda_function_name}"
  retention_in_days = 14
}

resource "aws_lambda_permission" "allow_bucket" {
  statement_id  = "AllowExecutionFromS3Bucket"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.stp_lambda.arn
  principal     = "s3.amazonaws.com"
  source_arn    = aws_s3_bucket.s3_for_image.arn
}
