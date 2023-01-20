# training golang

```
#setup environment
- git create a new repository on the command line
echo "# training-golang" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:jozemario/training-golang.git
git push -u origin main

- Golang Version Manager
$ curl -sSL https://raw.githubusercontent.com/voidint/g/master/install.sh | bash
$ echo "unalias g" >> ~/.bashrc # 可选。若其他程序（如'git'）使用了'g'作为别名。
$ source "$HOME/.g/env"

g ls-remote
g install 1.19.5
g use 1.19.5
 
```

### bfe
```
Download source code
$ git clone https://github.com/bfenetworks/bfe
Build
Execute the following command to build bfe:
$ cd bfe
$ make
!!! tip If you encounter an error such as "https fetch: Get ... connect: connection timed out", please set the GOPROXY and try again. See Installation FAQ

Execute the following command to run tests:
$ make test
Executable object file location:
$ file output/bin/bfe
output/bin/bfe: ELF 64-bit LSB executable, ...
Run
Run BFE with example configuration files:
$ cd output/bin/
$ ./bfe -c ../conf -l ../log
```

### go linters
```
Installing GolangCI-Lint
Use the command below to install golangci-lint locally on any operating system. Other OS-specific installation options can be found here.

$ go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
copy
Once installed, you should check the version that was installed:

$ golangci-lint version
golangci-lint has version v1.40.1 built from (unknown, mod sum: "h1:pBrCqt9BgI9LfGCTKRTSe1DfMjR6BkOPERPaXJYXA6Q=") on (unknown)
copy
You can also view the all the available linters through the following command:

$ golangci-lint help linters

If you run the enabled linters at the root of your project directory, you may see some errors. Each problem is reported with all the context you need to fix it including a short description of the issue, and the file and line number where it occurred.

$ golangci-lint run # equivalent of golangci-lint run ./...

You can also choose which directories and files to analyse by passing one or more directories or paths to files.

$ golangci-lint run dir1 dir2 dir3/main.go
```

### dev notes
```
go mod init unit-test  //create module
go mod tidy  //sync libraries
go run .  //run project
go test -bench=. 
go test -v 

go test -cover 
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out

---
Lint
golangci-lint help linters
golangci-lint run -v
golangci-lint run --fix

```