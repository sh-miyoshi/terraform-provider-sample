terraform {
  required_providers {
    sample = {
      version = "0.0.1"
      source = "github.com/sh-miyoshi/sample"
    }
  }
}

provider "sample" {
}

data "sample_storage" "storage1" {
  name = "storage1"
}

resource "sample_vm" "vm2" {
  name   = "vm2"
  cpu    = 1
  memory = 4096
  external_storage_id = data.sample_storage.storage1.id
}

