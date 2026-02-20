job "portfolio" {
  datacenters = ["eu-west"]
  namespace   = "app-portfolio"
  type        = "service"

  group "portfolio-group" {

    count = 1

    network {
      mode = "bridge"
    }

    service {
      name = "app-portfolio-service"
      port = "8080"

      tags = [
        "traefik.enable=true",
        "traefik.http.routers.portfolio.rule=Host(`conamu.com`)",
        "traefik.consulcatalog.connect=true"
      ]

      connect {
        sidecar_service {}
      }
    }

    task "portfolio-task" {
      driver = "docker"

      kill_timeout = "60s"
      kill_signal  = "SIGINT"

      config {
        image = "conamu470/portfolio"
      }
    }
  }
}
