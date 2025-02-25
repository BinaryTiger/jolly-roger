terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.0"
    }
  }
}

provider "docker" {}

resource "docker_image" "nats" {
  name = "nats:latest"
  keep_locally = true
}

resource "docker_container" "nats_server" {
  name  = "nats_test_server"
  image = docker_image.nats.image_id
  
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
