# IQM-assignment
## IQM-assignment-2
How to build an echo server from Dockerfile?
- 1. Open terminal and create a folder and download this branch in the folder: 
```Shell
$ mkdir mytmp
$ cd mytmp
$ git clone git@github.com:dalpengholic/IQM-assignment.git
```

- 2. Checkout to task-2 branch
`git checkout task-2`

- 3. Build docker image
`docker build -t my_golang_echo_server .`

- 4. Run echo server container
```Shell
// 1. Default command 
// Using default message "hello"
// and default port number "8000"
docker run --rm -p 8000:8000 my_golang_echo_server:latest

// 2. Advanced command
// Using user defined message  "-e EnvMessage=moimoi"
// and using user defined port number "-e EnvPort=8888"
docker run --rm -p 8888:8888 -e EnvPort=8888 -e EnvMessage=moimoi my_golang_echo_server:latest



```

## IQM-assignment-1
Decision Point-1: Which langauge would be better to implement simple echo server?
- Candidates: Golang, Python
- After a bit of searching for each language, I ended up choosing Golang.
- Compared to Python, with Golang it seemed easier to implement an Echo server through the default libraries without any web framework. Also, the number of lines of code required to write seemed short. 
