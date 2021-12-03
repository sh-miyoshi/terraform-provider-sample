terraform {
  required_providers {
    sample = {
      source = "github.com/sh-miyoshi/sample"
    }
  }
}

resource "sample_vm" "vm1" {
  name   = "vm1"
  cpu    = 1
  memory = 2048
}
