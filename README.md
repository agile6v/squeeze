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

<img src="https://github.com/agile6v/squeeze/blob/master/squeeze.jpg" width="60%" height="50%">


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

Squeeze provides two ways to start a task.

### 1. Squeeze Client Tool

Squeeze provide client-side tool that can interact with the squeeze cluster to start or cancel task and provide both synchronous and asynchronous modes. After the squeeze cluster is built, we can start a task. Squeeze supports a variety of protocols and each protocol need to provide different parameters. Use http protocol as the example.

First, Have a look at http protocol need to provide parameters.

```shell
$ squeeze client http --help

http protocol benchmark

Usage:
  squeeze client http [flags]

Flags:
  -d, --body string           Request body string
  -D, --bodyfile string       Request body from file
  -c, --concurrency int       Number of multiple requests to make at a time (default 1)
  -T, --content-type string   Content-type header to use for POST/PUT data (default "text/plain")
      --disable-compression   Disable compression of body received from the server.
      --disable-keepalive     Disable keepalive, connection will use keepalive by default.
  -z, --duration int          Duration of application to send requests. if duration is specified, n is ignored.
      --header strings        Custom HTTP header.(Repeatable)
  -h, --help                  help for http
      --http2                 Enable http2
      --maxResults int        The maximum number of response results that can be used (default 1000000)
  -m, --method string         Method name (default "GET")
  -x, --proxy string          HTTP Proxy address as host:port
  -q, --rateLimit int         Rate limit, in queries per second (QPS). Default is no rate limit
  -n, --requests int          Number of requests to perform (default 2147483647)
  -s, --timeout int           Seconds to max. wait for each response(Default is 30 seconds) (default 30)

Global Flags:
      --alsologtostderr                  log to standard error as well as files
      --callback string                  If this call is asynchronous then stress result will be sent to the address.
      --httpAddr string                  The address and port of the Squeeze master or slave. (default "http://127.0.0.1:9998")
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

Second, Initiate 100 requests to www.baidu.com using 10 concurrency.

```shell
$ ./squeeze client http http://www.baidu.com -n 100 -c 10

Summary:
  Requests:	100
  Total:	1.3025 secs
  Slowest:	1.0686 secs
  Fastest:	0.0089 secs
  Average:	0.1260 secs
  Requests/sec:	76.7757

  Total data:	8100 bytes
  Size/request:	81 bytes

Latency distribution:
  10% in 0.0104 secs
  25% in 0.0170 secs
  50% in 0.0228 secs
  75% in 0.0309 secs
  90% in 1.0503 secs
  95% in 1.0662 secs
  99% in 1.0686 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.1260 secs, 0.0089 secs, 1.0686 secs
  DNS-lookup:	0.1016 secs, 0.0000 secs, 1.0159 secs
  req write:	0.0001 secs, 0.0000 secs, 0.0027 secs
  resp wait:	0.0213 secs, 0.0000 secs, 0.0436 secs
  resp read:	0.0005 secs, 0.0000 secs, 0.0148 secs

Status code distribution:
  [200]	100 responses
```

### 2. Squeeze UI

still under development.

# Contributing
If you are interested in contributing to the Squeeze project, welcome to submit a PR.

# Future Works

We use [Project Board](https://github.com/agile6v/squeeze/projects/1) to track future works. If you are interested in the feature in the TODO list, please feel free to submit an issue.

# License

Squeeze is released under the Apache 2.0 license. See [LICENSE](https://github.com/agile6v/squeeze/blob/master/LICENSE) for details.

[architecutre]: https://github.com/agile6v/squeeze/blob/master/squeeze.jpg
