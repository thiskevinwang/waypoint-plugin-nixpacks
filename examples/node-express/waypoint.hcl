project = "express-nixpacks-example"
app "express-nixpacks-example" {
  build {
    use "nixpacks" {}
  }
  deploy {
    use "null" {}
  }
  release {
    use "null" {}
  }
}
