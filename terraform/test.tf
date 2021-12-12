terraform {
  required_providers {
    sample = {
      versions = ["0.0.1"]
      source = "github.com/sh-miyoshi/sample"
    }
  }
}

provider "sample" {
  app_url = "http://localhost:4567"
}

resource "sample_storage" "storage1" {
  name = "storage1"
  size = 50
}

resource "sample_vm" "vm1" {
  name   = "vm1"
  cpu    = 1
  memory = 2048
}
