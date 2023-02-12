project = "express-nixpacks-example"
app "express-nixpacks-example" {
  build {
    use "nixpacks" {
      source = "${path.app}/foo"
    }
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
