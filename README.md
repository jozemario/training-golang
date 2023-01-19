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

### dev notes
```
go mod init unit-test  
go mod tidy  
go run . 
go test -cover 
go test -bench=. 
go test -v 
```