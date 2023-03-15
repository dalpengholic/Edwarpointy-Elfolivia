# IQM-assignment

Main points of consideration during the whole assginment
```
- Provide simple and to-the-point solutions to the tasks.
- Choose the tooling which you consider minimal for the given assignment.
- Make sure to include instructions about how to reproduce your solution.
- Your MVP will not go straight into production. However, your programmer peers will look at
it, and it will be hard to convince them if the code is messy.
```

## IQM-assignment-1
### Decision Point-1: Which langauge would be better to implement simple echo server?
- Candidates: Golang, Python
- After a bit of searching for each language, I ended up choosing Golang.
- Compared to Python, with Golang it seemed easier to implement an Echo server through the default libraries without any web framework. Also, the number of lines of code required to write seemed short.

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

- 3. Run echo server
```Shell
## 1. Default command 
$ ./webserver

## 2. Advanced command
## Using user defined message  "-c moimoi"
## and using user defined port number "-p 8888"
$ ./webserver -c moimoi -p 8888
```

- 4. Open new terminal and run curl command
```
## In case of running server with the advanced command
$ curl http://localhost:8888
```
 
