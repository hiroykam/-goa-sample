# simple goa sample program for the trial usage

## environment
- docker
- goa
- dep
- gorm
- sql-migrate

## installation &setup

- install mercury  
run `brew install mercurial` to avoid the following hang of freeze during fetching of `goa`.

```$xslt
$ pstree 69752
-+= 69752 hiroyukikamisaka dep ensure -v
 \-+= 71086 hiroyukikamisaka /Applications/Xcode.app/Contents/Developer/usr/bin/git ls-remote ssh://git@bitbucket.org/pkg/inflect
   \--- 71087 hiroyukikamisaka /usr/bin/ssh git@bitbucket.org git-upload-pack '/pkg/inflect'
```

- install goa
run `go get -u github.com/goadesign/goa/...` to get goa and goagen

- install dep  
run `go get -u github.com/golang/dep/cmd/dep` to get dep

- launch docker
```
$ make read-env
$ make docker-build
$ make docker-run
```

only the purpose of restart, run `make docker-run`.


## Development

run following commands after update the go files in `design` directory.

```
$ make controller
$ make app
```

To update the swagger document, run `make swagger`.
