terraform {
  required_providers {
    pocketbase = {
      source = "hashicorp.com/natrontech/pocketbase"
    }
  }
}

provider "pocketbase" {
  endpoint = "http://127.0.0.1"
  identity = "admin@natron.io"
  password = "0123456789"
}
