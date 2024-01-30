package global

import "github.com/whoisnian/glb/config"

var CFG Config

type Config struct {
	Debug bool `flag:"d,false,Enable debug output"`

	DryRun  bool `flag:"dry-run,false,Perform a trial run without actual changes"`
	Version bool `flag:"v,false,Show version and quit"`

	URL    string `flag:"u,https://example.com,Custom http request url"`
	Method string `flag:"m,GET,Custom http request method"`
}

func SetupConfig() {
	err := config.FromCommandLine(&CFG)
	if err != nil {
		panic(err)
	}
}
