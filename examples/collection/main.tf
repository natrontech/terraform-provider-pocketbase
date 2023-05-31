terraform {
  required_providers {
    pocketbase = {
      source = "hashicorp.com/natrontech/pocketbase"
    }
  }
}

provider "pocketbase" {
  endpoint = "http://127.0.0.1:8090"
  identity = "admin@natron.io"
  password = "0123456789"
}

resource "pocketbase_collection" "example" {
  # name = "example"
  # type = "base"
  # schema = [
  #   {
  #     name     = "title"
  #     type     = "text"
  #     required = true
  #     # options = {
  #     #   min = 10
  #     # }
  #   },
  #   {
  #     name = "status"
  #     type = "bool"
  #   }
  # ]
}