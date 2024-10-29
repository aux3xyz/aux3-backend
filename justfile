default: list

build:
  go build .

clean:
  rm -f aux3.xyz

test:
  go test ./...

list:
  @just --list --unsorted

help:
  echo "use just commands to perform actions"
  @just list

alias b := build
alias t := test
alias l := list
alias h := help
