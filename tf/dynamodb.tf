resource "aws_dynamodb_table" "table_for_image" {
  name           = "${var.dynamodb.table_name_prefix}-${random_string.lower.result}"
  billing_mode   = var.dynamodb.billing_mode
  read_capacity  = var.dynamodb.read_capacity
  write_capacity = var.dynamodb.write_capacity
  hash_key       = var.dynamodb.hash_key
  attribute {
    name = var.dynamodb.hash_key
    type = var.dynamodb.hash_key_attr_type
  }
  tags = {
    Name = var.tag.tag_name
  }
}
