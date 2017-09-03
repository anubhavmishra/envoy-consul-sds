job "envoy-consul-sds" {
  datacenters = ["dc1"]
  type = "service"
  group "webserver" {
    count = 1
    task "envoy-consul-sds" {
      driver = "docker"
      config {
        image = "anubhavmishra/envoy-consul-sds:latest"
        network_mode = "host"
      }
      service {
        name = "envoy-consul-sds"
        tags = ["sds"]
        port = "webserver"
        check {
          type     = "tcp"
          port     = "webserver"
          interval = "10s"
          timeout  = "2s"
        }
      }
      resources {
        network {
          mbits = 10
          port "webserver" {
            static = 8080
          }
        }
      }
    }
  }
}