resource "aws_cloudfront_origin_access_identity" "origin_access_identity" {
  comment = "CloudFront identity for stp"
}

resource "aws_cloudfront_distribution" "cloudfront_for_image" {
  origin {
    domain_name = aws_s3_bucket.s3_for_image.bucket_regional_domain_name
    origin_id   = var.s3.origin_id

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.origin_access_identity.cloudfront_access_identity_path
    }
  }

  enabled             = var.cloudfront.enabled
  is_ipv6_enabled     = var.cloudfront.is_ipv6_enabled
  comment             = var.cloudfront.comment
  default_root_object = var.cloudfront.default_root_object

  logging_config {
    include_cookies = var.cloudfront.logging_config.include_cookies
    bucket          = aws_s3_bucket.log_bucket.bucket_domain_name
    prefix          = var.cloudfront.logging_config.prefix
  }

  aliases = var.cloudfront.aliases

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = var.s3.origin_id
    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }
    viewer_protocol_policy = var.cloudfront.viewer_protocol_policy
    min_ttl                = var.cloudfront.min_ttl
    default_ttl            = var.cloudfront.default_ttl
    max_ttl                = var.cloudfront.max_ttl
  }

  price_class = var.cloudfront.price_class
  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  tags = {
    Name = var.tag.tag_name
  }
  viewer_certificate {
    cloudfront_default_certificate = true
  }
}
