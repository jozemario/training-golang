#vault server mockup for test
```
go.mod
github.com/hashicorp/vault v1.11.1
	github.com/hashicorp/vault/api v1.7.2
	github.com/hashicorp/vault/sdk v0.5.3
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.8.0

go mod init unit-test  //create module
go mod tidy  //sync libraries
go build
go run .

go test -covermode=atomic -coverprofile=./coverage.txt -v --run TestReadSecret

#####
-log json Obj map[string]interface{}
b, _ := json.Marshal(jsonObj);fmt.Println(string(b))

```