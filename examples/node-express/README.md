# express-nixpacks-example

This is the source code for a barebones express.js server application, running on port `3000`.

## Quickstart

To test the image that is built via `waypoint` & `waypoint-plugin-nixpacks`, run the following:

```bash
docker run -it -p 3000:3000  thekevinwang/express-nixpacks-example
# Visit localhost:3000
# Ctrl+c to stop the container
```

## Running with Waypoint

> **Note**: The following steps assume you have the `waypoint` CLI installed and a Waypoint server running. Checkout [HCP Waypoint][hcp-waypoint] for if you want the ease of having a fully-managed Waypoint server.

First make sure a `waypoint-plugin-nixpacks` executable exists alongside `waypoint.hcl`

To build the application yourself and publish the resulting image to your own image registry, such as Docker Hub, ensure a `docker_auth.json` file exists, alongside `waypoint.hcl`

- See the `waypoint.hcl` file's `filebase64` usage as well as the `docker_auth.example.json` file to gain a better idea of how authenication is achieved.

Update the `image` field under `app => build => registry => use => docker`

Run the usual waypoint commands:

```bash
waypoint init
waypoint build
```

Head over to your image registry and check that a new image was uploaded.

[hcp-waypoint]: https://portal.cloud.hashicorp.com/
