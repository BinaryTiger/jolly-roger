terraform {
  required_providers {
    podman = {
      source  = "fbreckle/podman"
      version = "~> 0.5.0"
    }
  }
}

provider "podman" {}

resource "podman_image" "nats" {
  name = "nats:latest"
  keep_locally = true
}

resource "podman_container" "nats_server" {
  name  = "nats_test_server"
  image = podman_image.nats.name
  
  ports {
    internal = 4222
    external = 4222
  }
  
  ports {
    internal = 8222
    external = 8222
  }
  
  command = ["--config", "/nats-server.conf"]
  
  upload {
    content = file("${path.module}/../../nats.conf")
    file    = "/nats-server.conf"
  }
}

output "nats_uri" {
  value = "nats://localhost:4222"
}
