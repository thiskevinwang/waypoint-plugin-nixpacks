project = "express-nixpacks-example"
app "express-nixpacks-example" {
  build {
    use "nixpacks" {}
    registry {
      use "docker" {
        image = "thekevinwang/express-nixpacks-example"
        tag   = "latest"
        encoded_auth = filebase64("${path.app}/docker_auth.json")
      }
    }
  }
  deploy {
    use "null" {}
  }
  release {
    use "null" {}
  }
}
