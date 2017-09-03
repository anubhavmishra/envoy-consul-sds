job "consul" {
  datacenters = ["dc1"]
  type = "system"
  update {
    stagger = "5s"
    max_parallel = 1
  }
  group "consul-agent" {
    task "consul-agent" {
      driver = "exec"
      config {
        command = "consul"
        args = ["agent", "-data-dir", "/var/lib/consul"]
      }
      artifact {
        source = "https://releases.hashicorp.com/consul/0.9.2/consul_0.9.2_linux_amd64.zip"
        options {
          checksum = "sha256:0a2921fc7ca7e4702ef659996476310879e50aeeecb5a205adfdbe7bd8524013"
        }
      }
      resources {
        cpu = 500
        memory = 64
        network {
          mbits = 1
          port "server" {
            static = 8300
          }
          port "serf_lan" {
          }
            static = 8301
          }
          port "serf_wan" {
            static = 8302
          }
          port "rpc" {
            static = 8400
          }
          port "http" {
            static = 8500
          }
          port "dns" {
            static = 8600
          }
        }
      }
    }
  }
}