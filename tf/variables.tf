variable "dynamodb" {
  default = {
    table_name_prefix  = "stp"
    billing_mode       = "PROVISIONED"
    read_capacity      = 20
    write_capacity     = 20
    hash_key           = "FileName"
    hash_key_attr_type = "S"
  }
}

variable "s3" {
  default = {
    bucket_name_prefix  = "stp"
    origin_id           = "stp_s3_origin" // for setup cloudfront origin.
    origin_folder_path  = "org/" // for keeping original image file.
    resized_folder_path = "resize/" // for keeping resized image file.
  }
}

variable "lambda" {
  default = {
    function_name_prefix = "stp"
    handler              = "main"
    timeout              = 300
    memory_size          = 196
    environment = {
      VLOG    = "true"
      PREFIX  = "resized-"
      RESIZEX = 900
      RESIZEY = 900
    }
  }
}

variable "cloudfront" {
  default = {
    aliases             = []
    comment             = "cloudfront for stp"
    default_root_object = "index.html"
    logging_config = {
      prefix          = ""
      include_cookies = false
    }
    viewer_protocol_policy = "https-only"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
    price_class            = "PriceClass_200"
    enabled                = true
    is_ipv6_enabled        = true
  }
}


variable "tag" {
  default = {
    tag_name = "managed_by_stp"
  }
}
