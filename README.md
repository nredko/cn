# CTRL-T ![CI](https://github.com/codenotary/ctrlt/workflows/CI/badge.svg)

* `ctrlt` - a notarization daemon
* `cn` - a CLI tool for notarizing, un-trusting and verifying docker images

## Requirements

* Golang
* [immudb](https://github.com/codenotary/immudb)

## Building the Project
To build the project binaries, execute:

    make

## Running the Server
There is a standalone docker-compose configuration supplied that bootstraps
all required components. Run it via:

    make run

If you want to run CTRL-T in standalone mode, start `immudb` and then execute

    IMMU_ADDR=$address IMMU_PORT=$port ./ctrlt

and follow the instructions on the CLI.

## Notarizing
Once `immudb` and `ctrlt` are up and running, you can start notarizing docker
images via:

    docker run -it alpine

    ./cn notarize alpine
    ./cn verify alpine
    ./cn list

The result is also reflected in the web UI running at http://localhost:4040.
