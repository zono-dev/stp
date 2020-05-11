data "aws_iam_policy_document" "s3_policy_for_stp" {
  statement {
    actions = ["s3:GetObject"]
    resources = [
      "${aws_s3_bucket.s3_for_image.arn}/*"
    ]
    principals {
      type        = "AWS"
      identifiers = ["${aws_cloudfront_origin_access_identity.origin_access_identity.iam_arn}"]
    }
  }
}

data "aws_iam_policy_document" "policy_doc_for_stp" {
  // for accessing to s3 bucket
  statement {
    actions = [
      "s3:GetObject",
      "s3:PutObject",
    ]

    resources = [
      "${aws_s3_bucket.s3_for_image.arn}/*"
    ]
  }

  // for accessing to dynamodb table
  statement {
    actions = [
      "dynamodb:DeleteItem",
      "dynamodb:GetItem",
      "dynamodb:PutItem",
      "dynamodb:UpdateItem",
    ]

    resources = [
      aws_dynamodb_table.table_for_image.arn
    ]
  }

  // AWSLambdaBasicExecutionRole
  statement {
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = [
      "*"
    ]
  }
}

// making policy for stp working
resource "aws_iam_policy" "policy_for_stp" {
  name   = "policy_for_stp"
  path   = "/"
  policy = data.aws_iam_policy_document.policy_doc_for_stp.json
}

data "aws_iam_policy_document" "assume_role_for_stp" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "role_for_stp" {
  name               = "role_for_stp"
  assume_role_policy = data.aws_iam_policy_document.assume_role_for_stp.json
}

// attach role with policy
resource "aws_iam_role_policy_attachment" "attachment_for_stp" {
  role       = aws_iam_role.role_for_stp.name
  policy_arn = aws_iam_policy.policy_for_stp.arn
}

