# go-templates
Start a new project from templates with [gonew](https://go.dev/blog/gonew).

## example
Start a new project `mycli` from `go-templates/cli`
```sh
## 1.install gonew
go install golang.org/x/tools/cmd/gonew@latest

## 2.create new project
gonew github.com/whoisnian/go-templates/cli github.com/whoisnian/mycli
# gonew: initialized github.com/whoisnian/mycli in ./mycli

## 3.push to github
cd ./mycli
git init
git add .
git commit -m "ADD: gonew from github.com/whoisnian/go-templates/cli"
git remote add origin git@github.com:whoisnian/mycli.git
git push -u origin master
```

## templates
* [cli](https://github.com/whoisnian/go-templates/tree/master/cli): A command line tool that sends a request and prints the response.
* [server](https://github.com/whoisnian/go-templates/tree/master/server): An HTTP server that serves plain messages CRUD.
