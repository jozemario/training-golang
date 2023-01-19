### unit-testing
```
we are going to take a look unit testing in Go. As you know Go is a programming language designed to develop software services by Google engineers who name is Rob Pike, Robert Griesemer and Ken Thompson.

If you are not familiar The Go Programming Language check this document The Golang Documentation

Introduction
Unit test is a unit function that tests your codes, components and modules. The purpose of unit test is clear of bugs from code, increase stability of code, provides correctness when changing code. The unit test is the first level of software test and followed by integration tests and ui tests.


Test Pyramid
In this tutorial we will create a small program and write test-friendly code using Go Testing package and go test command. While we write test-friendly code, we will implement different approaches which is Table-Driven Testing, Test Coverage, Benchmark and Document Example.

Prerequisites
Familiar with Go Programming Language.
Installing latest Go version on your local machine. (You can follow the instructions How to Install Go on Windows, MacOS, Linux)
IDE for write code and run tests (like VS Code, Goland, SublimeText …)
Create Go Project
In your local machine create directory for your project called unit-test from terminal

$ mkdir unit-test
Open project folder location with preferred IDE

If you are using VS Code, you can open folder with following command

$ cd unit-test
$ code .
Create main.go file. The main.go file is entry point for Golang Programming language. Then run following command to init your module.

$ go mod init unit-test
Finally we are ready to implement unit test with Go. Your project folder looks like below

.
├── go.mod
└── main.go
Before start the coding, we should know about Go test package (testing.T) and go testcommand. Go’s standart library provides the built-in testing package and go test command.

testing.T provides to implement unit test
go test command provides run the tests
Basic unit test looks like in Go.


First parameter must be t *testing.T
Test function name should be begin with Test prefix and following by a word or phrase
Test file name should be end with _test.go suffix.
t.Error() provides failure information
t.Log() provides non-failing information
Write Unit Test
Create greeting.go file in your project and add the following code


The greeting.go file provides sayHello function. The purpose of the function accept one argument which is name. If the name argument is empty returns “Hello Anonymous” otherwise returns name with “Hello” prefix.

Now we can test sayHello function. Create greeting_test.go file in your project and add the following code.


We have Test_SayHello_Valid_Argument unit test. It can be many tests in a single file. In the test function, we pass the valid argument and check if the result as expected. If the result is not valid, we call the t.Errorf()indicate that there was an error. You can call t.Logf()non-debugging information if you want to log test results.

It is time to execute test function. Run the following command

$ go test
You will receive following results

PASS
ok      unit-test       0.465s
PASS means the code is working as expected.

You can run unit test with more information using verbose with -v flag.

$ go test -v
The results will be like following.

=== RUN   Test_SayHello_ValidArgument
    greeting_test.go:16: "sayHello('Mert')" SUCCEDED, expected -> Hello Mert, got -> Hello Mert
--- PASS: Test_SayHello_ValidArgument (0.00s)
PASS
ok      unit-test       0.326s
So far so good, the test is passed. What if the test fails? Now we can change the sayHello function to fail test. Change the ‘Hello’ with ‘Hola’ in sayHello function then run test again and see the results what is happening.

Modified sayHello function like this.


Run the following command

$ go test -v
You will receive the results like following

=== RUN   Test_SayHello_ValidArgument
    greeting_test.go:14: "sayHello('Mert')" FAILED, expected -> Hello Mert, got -> Hola Mert
--- FAIL: Test_SayHello_ValidArgument (0.00s)
FAIL
exit status 1
FAIL    unit-test       0.611s
As we mentioned that we have many tests function in a single file. You can run a specific test from many tests with -run flag.

Add sayGoodBye function in greeting.go file. The sayGoodBye function looks like following.


Add sayGoodBye’s test into greeting_test.go


Now we have two unit test in a single file. When you run the tests with go test -v command

$ go test -v
You will receive the result of two tests like following

=== RUN   Test_SayHello_ValidArgument
    greeting_test.go:16: "sayHello('Mert')" SUCCEDED, expected -> Hello Mert, got -> Hello Mert
--- PASS: Test_SayHello_ValidArgument (0.00s)
=== RUN   Test_SayGoodBye
    greeting_test.go:28: "sayGoodBye('Mert')" SUCCEDED, expected -> Bye Bye Mert, got -> Bye Bye Mert
--- PASS: Test_SayGoodBye (0.00s)
PASS
ok      unit-test       0.219s
So far so good, now we can run specific test with -run flag. -run flag takes matching test name.

$ go test -v -run=Test_SayGoodBye
The result will be like following

=== RUN   Test_SayGoodBye
    greeting_test.go:28: "sayGoodBye('Mert')" SUCCEDED, expected -> Bye Bye Mert, got -> Bye Bye Mert
--- PASS: Test_SayGoodBye (0.00s)
PASS
ok      unit-test       0.475s
Table-Driven Test
You want to test your code with many inputs and expected result of these inputs. The best approach is creating an array for inputs and run tests each item in array and get expected result.

Let’s modify Test_SayHello_ValidArgument function for Table-Driven approach.


Run the test with -v flag

$ go test -v -run=Test_SayHello_ValidArgument
You will receive the results like following.

=== RUN   Test_SayHello_ValidArgument
    greeting_test.go:36: "sayHello('Yemeksepeti')" succeded, expected -> Hello Yemeksepeti, got -> Hello Yemeksepeti
    greeting_test.go:36: "sayHello('Banabi')" succeded, expected -> Hello Banabi, got -> Hello Banabi
    greeting_test.go:36: "sayHello('Yemek')" succeded, expected -> Hello Yemek, got -> Hello Yemek
--- PASS: Test_SayHello_ValidArgument (0.00s)
PASS
ok      unit-test       0.288s
Test Coverage
Test coverage is measurement percentage of code coverage in your application. It’s important to know how much of your code the tests cover. In this way, you can see which parts of the code you have tested and which parts we have not tested.

Go’s standart library provides the built-in test cover to check your code coverage

Run the following command

$ go test -cover
You will receive the following results

PASS
coverage: 66.7% of statements
ok      unit-test       0.585s
Our test passed but coverage result is only %66.7. It means %66.7 of your code executed by the test. At this point we need to know what we missed to cover in the test.

go test command has -coverprofile flag. It allows you to export the test coverage results to a file. -coverprofile flag takes an argument that filename to output

run the following command and export cover profile to a file.

$ go test -coverprofile=cover_out
You receive the following result and exported cover_out file in your project directory

PASS
coverage: 66.7% of statements
ok      unit-test       0.703s
Your project folder looks like following and you can see the cover_out file in your project.

.
├── cover_out
├── go.mod
├── greeting.go
├── greeting_test.go
└── main.go
Extracted cover_outfile does not mean anything by self. Go provides cover tool command to convert cover profile to HTML file. Through this you can present code coverage result in a web browser

Run the following command

$ go tool cover -html=cover_out -o cover_out.html
You can see go tool command extracted cover_out.html file based on cover_out file. Open the cover_out.html file in a web browser and see the following result


The code in red color is not covered by the test. As we see above, we don’t handle test of sayHello when we passed empty argument to name. Thus you can see which part of the test is missing and increase your result of the test coverage accordingly.

Benchmarks with Go
With Benchmarking you can measure performance of the code and see the impact of the changes you make the code so you can optimized your source code.

The file name must be begin with Benchmark prefix as unit test file name convention. (BenchmarkSomething(*testing.B))

Add BenchmarkSayHello function in to greeting_test.go file.


Benchmark function must be take argument testing.B
Benchmark must run N times the testing.B provides (N is an integer type and adjusted from Go)
Executed by the “go test” command when its -bench flag is provided.
The --bench flag accepts its arguments in the form of a regular expression.
Run the following command

$ go test -bench=.
If you have many Benchmark functions in a single file you can declare function name explicitly like following command

go test -bench=BechmarkSayHello
You will receive the results like following

goos: darwin
goarch: amd64
pkg: unit-test
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSayHello-12            11307939                92.46 ns/op
PASS
ok      unit-test       1.920s
The Benchmark result means benchmark is passed. The loop ran 11.307.939 times and per loop speed is 92.46 nanosecond.

Documenting Go Code With Example
You can documenting Go codes with Example approach. Go is focused on documenting and example codes add different dimension to documenting and testing. Example approach based on existing method or function. It should show to users how to use your codes.

Add Example approach for sayHello function like following.


Example approach must start with Example prefix and following by existing function name
fmt package is import to list what you expected and match the output
Output: is document the expected output
Run the following command

$ go test -v
You will receive results and you can see Example approach was executed by the test.

=== RUN   Test_SayHello_ValidArgument
    greeting_test.go:16: "sayHello('Mert')" SUCCEDED, expected -> Hello Mert, got -> Hello Mert
--- PASS: Test_SayHello_ValidArgument (0.00s)
=== RUN   Test_SayGoodBye
    greeting_test.go:49: "sayGoodBye('Mert')" SUCCEDED, expected -> Bye Bye Mert, got -> Bye Bye Mert
--- PASS: Test_SayGoodBye (0.00s)
=== RUN   ExampleSayHello
--- PASS: ExampleSayHello (0.00s)
This feature improves your documentation and also makes your unit test more power.

Conclusion
In this tutorial, you created a simple Go project and implemented unit test based on Table-Driven, Test Coverage, Benchmark and Example approaches.

Unit test is important because when you have written the code, you want work to as expected result even the code changes. By this mean, it will improve your confidence.

As we mentioned that Go’s standart library provides built-in testing package. It has different unit-test functionality. To learn more information check the Go official document
```