# AWS Copilot for LocalStack

This is a fork of the [AWS Copilot CLI](https://github.com/aws/copilot-cli) for use with [LocalStack](https://localstack.cloud).

## Details

This repo provides `copilotlocal`, a command-line interface (CLI) with the same features as the original `copilot` CLI, but using the local API endpoints provided by LocalStack. The patch applied in this repo essentially redirects any AWS API calls to the local endpoints under `http://localhost:4566`.

## Configurations

You can configure the following environment variables:
* `LOCALSTACK_HOSTNAME`: Target hostname under which LocalStack endpoints are available (default: `localhost`)
* `EDGE_PORT`: Target port under which LocalStack endpoints are available (default: `4566`)
* `LOCALSTACK_DISABLE`: Optional flag to disable the local endpoints and use the default behavior of `copilot` (set to `1` or `true` to enable)

## Downloading Releases

We provide pre-built binaries for MacOS/Linux/Windows. Please refer to the [releases page](./releases) to download the right binary for your operating system and CPU architecture.

### Creating a new release

The `mainline` branch should be up-to-date with upstream `mainline`, and the changes in the fork are stored on the `localstack` branch. Note that there is also a `master` branch, but that one is a bit behind, so we use `mainline` to leverage the latest changes. As a first step, we need to pull changes from upstream `mainline`, then rebase `localstack` onto `mainline` (this may require resolving some conflicts).

Then, install the `npm` dependencies in the `cf-custom-resources` folder:
```
cd cf-custom-resources; npm install
```

Then we can build the binaries by compiling the Go program for different operating systems / CPU architectures:
```
make release
```

Next, we look up the latest tag in the upstream repo (e.g., `v1.24.0`), and then create a tag with the same name on the `localstack` branch on our fork. Assuming the remote is named `localstack`, it can then be pushed to our fork via:
```
# delete any existing tag with that version number
git tag -d v1.24.0
# create a new tag against the HEAD of `localstack` branch
git tag v1.24.0
# push the branch to the fork repo
git push localstack v1.24.0
```

Next, we create a new release in the Github repo using the version tag we just created (e.g., `v1.24.0`). Make sure to upload the binary files generated in `bin/local/...` as part of the release assets.

## License

This library is licensed under the Apache 2.0 License.
