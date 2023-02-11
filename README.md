![Screenshot of waypoint-plugin-nixpacks in HCP Waypoint](https://user-images.githubusercontent.com/26389321/218119583-9dabe325-9295-4098-90ac-b9f32ccb7b74.png)

---

# waypoint-plugin-nixpacks

This [Waypoint][waypoint] plugin builds OCI images using [`nixpacks`][nixpacks].

This plugin only implements the [builder][builder] interface.

## Building

To build this plugin, run `make`. (See [`Makefile`](./Makefile) for more details.)
A plugin binary will be built and outputted to the [`./bin`](./bin) folder as well as copied into the example project folder.

This requires the following packages:

- `go`
- `protoc`
- `protoc-gen-go`

## Usage

The Waypoint runner (such as the CLI) that ends up picking up a job (such as `waypoint build`) will need [`nixpacks`][nixpacks] installed.

Check out the [examples](./examples) folder for usage examples.

[builder]: https://developer.hashicorp.com/waypoint/docs/extending-waypoint/plugin-interfaces/builder
[waypoint]: https://github.com/hashicorp/waypoint
[nixpacks]: https://github.com/railwayapp/nixpacks
