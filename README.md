# local-sns-go

A local (mock) version of AWS SNS

Currently supports just one endpoint - Publish.


## Install

```sh
go get -u github.com/NSmithUK/local-sns-go

cd $GOPATH/src/github.com/NSmithUK/local-sns-go

dep ensure

go install

```

## Run

```sh
$GOPATH/bin/local-sns-go

```
