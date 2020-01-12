# cob-example

This repository is an example of [cob](https://github.com/knqyf263/cob/releases).

## GitHub Actions

```
name: Bench
on: [push, pull_request]
jobs:
  test:
    name: Bench
    runs-on: ubuntu-latest
    steps:

    - name: Set up Go 1.13
      uses: actions/setup-go@v1
      with:
        go-version: 1.13
      id: go

    - name: Check out code into the Go module directory
      uses: actions/checkout@v1

    - name: Install GolangCI-Lint
      run: curl -sfL https://raw.githubusercontent.com/knqyf263/cob/master/install.sh | sudo sh -s -- -b /usr/local/bin

    - name: Run Benchmark
      run: cob -benchmem ./...
```

## Travis CI

```
dist: bionic
language: go
go:
  - 1.13.x

before_script:
  - curl -sfL https://raw.githubusercontent.com/knqyf263/cob/master/install.sh | sudo sh -s -- -b /usr/local/bin

script:
  - cob -benchmem ./...
```

## CircleCI

```
version: 2
jobs:
  bench:
    docker:
      - image: circleci/golang:1.13
    steps:
      - checkout
      - run:
          name: Install cob
          command: curl -sfL https://raw.githubusercontent.com/knqyf263/cob/master/install.sh | sudo sh -s -- -b /usr/local/bin
      - run:
          name: Run cob
          command: cob -benchmem ./...
workflows:
  version: 2
  build-workflow:
    jobs:
      - bench
```
