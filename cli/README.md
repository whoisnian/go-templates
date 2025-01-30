# go-templates/cli
A command line tool that sends a request and prints the response.

## usage
* Build binary: `./build/build.sh .`
* Show help: `./output/cli -help`
* Run with command-line arguments: `./output/cli -d -u https://ip.whoisnian.com`
* Run with environment variables: `CFG_DEBUG=true CFG_URL=https://ip.whoisnian.com ./output/cli`

## features
* [x] Parsing command-line arguments and environment variables
* [x] Logging to stderr with custom log level
* [x] Build script for multiple platforms
* [x] GitHub Actions workflow for tagged release (need `Read and write permissions` in `Settings > Actions > General > Workflow permissions`)
