provider "aws" {
  region                      = "sa-east-1"
  access_key                  = "tynamo"
  secret_key                  = "tynamo"
  skip_credentials_validation = true
  skip_requesting_account_id  = true
  skip_metadata_api_check     = true
}
