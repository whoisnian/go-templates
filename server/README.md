# go-templates/server
An HTTP server that serves plain messages CRUD.

## usage
* Build binary: `./build/build.sh .`
* Show help: `./output/server -h`
* Run with command-line arguments: `./output/server -l 0.0.0.0:9000`
* Run with environment variables: `CFG_LISTENADDR=0.0.0.0:9000 ./output/server`
* Run with config file: `./output/server -config ./config.json`

## features
* [x] Parsing command-line arguments, environment variables and config file
* [x] Logging as json lines to stderr with custom log level
* [x] Build script for multiple platforms
* [x] GitHub Actions workflow for tagged release (need `Read and write permissions` in `Settings > Actions > General > Workflow permissions`)
* [ ] Systemd unit configuration file
* [x] Docker images for multiple platforms
