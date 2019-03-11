# Squeeze

Squeeze is a modern, easy-to-use, and highly capable load-testing tool.  It uses the Master-Slave pattern to simulate any number of users hitting the target.  In addition, Squeeze provides the command line and web-based tool to create test tasks and display test results.

# Table of Contents
- [Features](#features)
- [Project Status](#project-status)
- [Architecture](#architecture)
- [Documentation](#documentation)
- [Prerequisite](#prerequisite)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [Future Works](#future-works)
- [License](#license)

# Features
* **Multiple Protocols**: HTTP1.0, HTTP1.1, HTTP2.0, HTTPS, Websocket

# Project Status
***Experimental***. This project is still under development.


# Architecture


# Documentation

# Prerequisite
* Requires Go 1.11 or higher
* Requires [protoc](https://github.com/google/protobuf)
* Requires [protoc-gen-go](https://github.com/google/protobuf)

# Installation
### 1. From source

Squeeze use the golang module mechanism to manage dependencies. We just need to type `make build`,  it will automatically download dependencies and generate binary. 

If you are a mainland China user, you may compile binary using golang vendor mechanism, type `make build-vendor`.

### 2. From a pre-built binary

We provide release binary for three platforms: Linux, OSX and Windows. You can download an appropriate release binary for your operating system. The list of binary releases is available for download from the [Release Pages](https://github.com/agile6v/squeeze/releases). If the binary doesn't work for you, you'll have to build from source.

### 3. With Docker or Kubernetes
Squeeze provides a corresponding docker image for each version and hosted on [Docker Hub](https://hub.docker.com/r/agile6v/squeeze). Therefore we can use docker-compose to quickly build a squeeze cluster to experience all the features.  Deployment on kubernetes will also support in the near future.

**Starting squeeze cluster with docker-compose:**

```shell
$ docker-compose up -d
```

**Scaling the slave with docker-compose:**

```shell
$ docker-compose up -d --scale slave=3
```

**Stoping squeeze cluster with docker-compose:**

```shell
$ docker-compose down
```

**Deploying on Kubernetes:**

```shell

```



# Usage

# Contributing
If you are interested in contributing to the Squeeze project, welcome to submit a PR.

# Future Works

We use [Project Board](https://github.com/agile6v/squeeze/projects/1) to track future works. If you are interested in the feature in the TODO list, please feel free to submit an issue.

# License

Squeeze is released under the Apache 2.0 license. See [LICENSE](https://github.com/agile6v/squeeze/blob/master/LICENSE) for details.
