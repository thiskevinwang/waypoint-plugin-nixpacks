project = "express-nixpacks-example"
app "express-nixpacks-example" {
  build {
    use "nixpacks" {}
    registry {
      use "docker" {
        image = "thekevinwang/express-nixpacks-example"
        tag   = "latest"
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
