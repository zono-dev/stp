resource "random_string" "lower" {
  length  = 16
  upper   = false
  lower   = true
  number  = false
  special = false
}