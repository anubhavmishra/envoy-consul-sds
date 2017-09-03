job "proxy" {
  datacenters = ["dc1"]
  type = "system"
  group "proxy" {
    count = 1
    task "envoy" {
      driver = "docker"
      config {
        image = "lyft/envoy:4640fc028d65a6e2ee18858ebefcaeed24dffa81"
        command = "/usr/local/bin/envoy"
        args = [
            "--concurrency 4",
            "--config-path /etc/envoy.json",
            "--mode serve",
        ]
        volumes = ["new/envoy.json:/etc/envoy.json" ]
        network_mode = "host"
      }
      artifact {
        source = "https://gist.githubusercontent.com/anubhavmishra/afe699320bdc4d855d13e7cc244822e0/raw/9cdd1f16355b1b7b0340c745ca650dc4a8fe0937/envoy.json"
      }
      template {
        source        = "local/envoy.json"
        destination   = "new/envoy.json"
        change_mode   = "restart"
      }
      resources {
        network {
          mbits = 10
          port "envoy" {
            static = 1010
          }
        }
      }
    }
  }
}