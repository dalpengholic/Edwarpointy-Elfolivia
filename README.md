## IQM-assignment-1
### Programming language decision: Golang 
When I had to choose a programming language for building a simple echo server, I had two options - Golang and Python. After some research, I chose Golang because it was easier to implement the server using its default libraries without any web framework, and the amount of code needed looked less than in Python.


### How to run
- 1. Open terminal and create a folder and download this branch in the folder: 
```Shell
$ mkdir mytmp
$ cd mytmp
$ git clone git@github.com:dalpengholic/IQM-assignment.git
```

- 2. Checkout to task-2 branch
```Shell
$ cd IQM-assignment
$ git checkout task-1
```

- 3. Build from source
```Shell
$ go build 
```

- 4. Run echo server
```Shell
## 1. Default command 
$ ./webserver

## 2. Advanced command
## Using user defined message  "-c moimoi"
## and using user defined port number "-p 8888"
$ ./webserver -c moimoi -p 8888
```

- 5. Open new terminal and run curl command
```
## In case of running server with the advanced command
$ curl http://localhost:8888
```
 
