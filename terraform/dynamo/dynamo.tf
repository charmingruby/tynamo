resource "aws_dynamodb_table" "examples_table" {
  name         = "examples"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "ID"

  attribute {
    name = "ID"
    type = "S"
  }
}
