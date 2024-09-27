terraform {
  # backend "s3" {
  #   bucket = "uryoya-terraformbook-up-and-running-state"
  #   # key: Terraformステートが書き込まれるべきS3バケット内のファイルパス
  #   key            = "global/s3/terraform.tfstate"
  #   region         = "us-east-2"
  #   dynamodb_table = "terraform-up-and-running-locks"
  #   # encrypt: この値をtrueに設定することでTerraformステートがS3に保存される際に暗号化される
  #   encrypt = true
  # }
}

provider "aws" {
  region = "us-east-2"
}

# Terraformステートを管理するためのS3バケット
resource "aws_s3_bucket" "terraform_state" {
  bucket = "uryoya-terraformbook-up-and-running-state"

  # 誤ってS3バケットを削除するのを防止
  # lifecycle {
  #   prevent_destroy = true
  # }
}

# ステートファイルの完全な履歴が見られるように、バージョニングを有効化
resource "aws_s3_bucket_versioning" "enabled" {
  bucket = aws_s3_bucket.terraform_state.id
  versioning_configuration {
    status = "Enabled"
  }
}

# デフォルトでサーバーサイド暗号化を有効化
resource "aws_s3_bucket_server_side_encryption_configuration" "default" {
  bucket = aws_s3_bucket.terraform_state.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

# 明示的にこのS3バケットに対する全パブリックアクセスをブロック
resource "aws_s3_bucket_public_access_block" "publick_access" {
  bucket                  = aws_s3_bucket.terraform_state.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# Terraformステートのロックを管理するDynamoDB
resource "aws_dynamodb_table" "terraform_locks" {
  name         = "terraform-up-and-running-locks"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}
