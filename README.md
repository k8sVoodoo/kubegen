# kubegen

`kubegen` is a CLI tool for generating, applying, and updating Kubernetes configurations. It simplifies the management of Kubernetes resources by automating common tasks.

## Features

- **Generate**: Create Kubernetes configuration files from templates.
- **Apply**: Apply configuration files to your Kubernetes cluster.
- **Update Image**: Update the container image in existing Kubernetes resources.

## Installation

1. **Build from Source**

   Clone the repository and build the tool:

   ```sh
   git clone <repository-url>
   cd kubegen
   go build -o kubegen

Add to PATH

Add the kubegen binary to your PATH for easier access.

## Usage
### Generate
Generate a Kubernetes configuration file.

```sh
kubegen generate --type <resource-type> --name <resource-name> [--namespace <namespace>] [--image <image>]
```
Flags:

--type, -t: Type of Kubernetes resource (e.g., deployment, statefulset, service).
--name, -n: Name of the Kubernetes resource.
--namespace, -N: Namespace for the Kubernetes resource (default: default).
--image, -i: Container image to use (default: my-image:latest).
Examples:

```sh
kubegen generate --type <resource-type> --name <resource-name> [--namespace <namespace>] [--image <image>]
```

### Apply
Apply a Kubernetes configuration file to the cluster.

```sh
kubegen apply <file-path>
```

Examples:

```sh
kubegen apply path/to/generated-deployment-my-deployment.yaml
```

### Update Image
Update the container image in a Kubernetes resource.

```sh
kubegen update-image --type <resource-type> --name <resource-name> --version <image-version>
```

Flags:

--type, -t: Type of Kubernetes resource (deployment, daemonset, statefulset).
--name, -n: Name of the Kubernetes resource.
--version, -v: New container image version (e.g., my-image:new-tag).

Examples:

```sh
kubegen update-image -t deployment -n my-deployment -v my-image:new-tag
```

### Help
For more information on each command, use the --help flag:

```sh
kubegen --help
kubegen generate --help
kubegen apply --help
kubegen update-image --help
```
