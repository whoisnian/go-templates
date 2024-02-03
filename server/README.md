# go-templates/server
An HTTP server that serves plain messages CRUD.

## usage
* Build binary: `./build/build.sh .`
* Show help: `./output/server -h`
* Run with command-line arguments: `./output/server -l 0.0.0.0:9000`
* Run with environment variables: `CFG_LISTENADDR=0.0.0.0:9000 ./output/server`

## features
* [ ] Parsing command-line arguments and environment variables
* [ ] Logging as json lines to stderr with custom log level
* [ ] Build script for multiple platforms
* [ ] GitHub Actions workflow for tagged release (need `Read and write permissions` in `Settings > Actions > General > Workflow permissions`)
* [ ] Systemd unit configuration file
* [ ] Docker images for multiple platforms
