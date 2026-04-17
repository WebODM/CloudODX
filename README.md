# CloudODX

![cloudodx-64x64](https://user-images.githubusercontent.com/1951843/51078515-02348000-1684-11e9-8f96-ed056b0cbe98.png)

A command line tool to process aerial imagery in the cloud via [NodeODX](https://github.com/WebODM/NodeODX)'s API.

## Getting Started

1. [Download the application](https://github.com/WebODM/CloudODX/releases) for Windows, Mac or Linux.
2. Extract the application in a folder of your choice (for example, `c:\odx`).
3. Open a command prompt and navigate to the folder (open the "Command Prompt" application, then `cd \odx`).
4. Run `odx c:\path\to\images --dsm`.

This command will process all the images in the directory `c:\path\to\images` and save the results (including an orthophoto, a point cloud, a 3D model and a digital surface model) to `c:\odx\output`. You can pass more options for processing by appending them at the end of the command. To see a list of options, simply issue:

`odx args`

See `odx --help` for more options.

## Using GCPs

To include a GCP for additional georeferencing accuracy, simply create a .txt file according to the [Ground Control Points format specification](https://docs.webodm.org/ground-control-points/#gcp-file-format) and place it along with the images.

## Processing Node Management

By default CloudODX will randomly choose a default node from the list of [publicly available nodes](https://github.com/WebODM/CloudODX/blob/master/public_nodes.json). If you are running your own processing node via [NodeODX](https://github.com/WebODM/NodeODX) you can add a node by running the following:

`odx node add mynode http://address:port`

Then run odx as following:

`odx -n mynode c:\path\to\images`

If no node is specified, the `default` node is selected. To see a list of nodes you can run:

`odx node -v`

For more information run `odx node --help`.

If you are interested in adding your node to the list of [public nodes](https://github.com/WebODM/CloudODX/blob/master/public_nodes.json) please open an [issue](https://github.com/WebODM/CloudODX/issues).

## Running From Sources

```bash
go get -u github.com/WebODM/CloudODX
go run github.com/WebODM/CloudODX/cmd/odx --help
```

## Building From Sources

We use [Goreleaser](https://goreleaser.com/) to build and deploy CloudODX. See Goreleaser's [documentation](https://goreleaser.com/) for installation and deployment instructions. You should just need to install the `goreleaser` application and then run:

`goreleaser release --skip-publish --snapshot`

## Reporting Issues / Feature Requests / Feedback

Please open an [issue](https://github.com/WebODM/CloudODX).

## Support the Project

There are many ways to contribute back to the project:

- ⭐️ us on GitHub.
- Help us test the application.
- Spread the word about WebODM on social media.
- Join a [community](https://webodm.org/community)
- Become a contributor!
