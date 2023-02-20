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

Also checkout https://github.com/thiskevinwang/waypoint-github-actions-test/blob/main/.github/workflows/workflow_dispatch.yml for full end to end usage in GitHub actions.

- This linked workflow uses a GitHub runner to connect to the Waypoint server on HCP, build `waypoint-plugin-nixpacks`, build the example app, and publish the resulting Docker image to DockerHub.

[builder]: https://developer.hashicorp.com/waypoint/docs/extending-waypoint/plugin-interfaces/builder
[waypoint]: https://github.com/hashicorp/waypoint
[nixpacks]: https://github.com/railwayapp/nixpacks

### Note on `nixpacks`

The machine (or runner) that executes this plugin will require `nixpacks` to be installed. Here are 2 different scenarios to help explain this.

### `waypoint up -local=true`

In this scenario, the `waypoint` CLI will serve as the **runner**, and maybe it is connecting out to the **server** running on HCP.

In this scenario the host machine executing `waypoint up -local=true` will also need `nixpacks` installed for the plugin to function correctly. This may already be the case if you're like me and doing development and testing from the same machine.

### `waypoint up -local=false`

In this scenario, the `waypoint` CLI only queues up jobs and a remote **runner** will be hosted elsewhere. That runner could be running on various platforms, such as:

- Directly on EC2, via
  ```bash
  waypoint runner agent
  ```
- As `docker` container, via
  ```
  waypoint runner install \
    -platform=docker \
    -server-addr=api.hashicorp.cloud:443 \
    -docker-runner-image=hashicorp/waypoint
  ```

In the EC2 option, the EC2 instance itself will need `nixpacks` installed.

In the Docker option, the `-docker-runner-image` will need `nixpacks` installed.
