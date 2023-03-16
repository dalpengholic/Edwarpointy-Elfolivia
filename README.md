## IQM-assignment-2
### Main concept of Dockerfile
To create a smaller Minimum-Viable Product (MVP) image, I used a two-stage approach called multi-stage builds. In the first stage, I only copied the necessary files to build the binary file, and in the second stage, I added those files to the busybox image. As a result, the MVP image size decreased from 314MB to 11.5MB.

### How to build an echo server from Dockerfile?
- 1. Open terminal and create a folder and download this branch in the folder: 
```Shell
$ mkdir mytmp
$ cd mytmp
$ git clone git@github.com:dalpengholic/IQM-assignment.git
```

- 2. Checkout to task-2 branch
```Shell
$ cd IQM-assignment
$ git checkout task-2
```
- 3. Build docker image
```Shell
$ docker build -t my_golang_echo_server .
```

- 4. Run echo server container
```Shell
// 1. Default command 
// Using default message "hello"
// and default port number "8000"
$ docker run --rm -p 8000:8000 my_golang_echo_server:latest

// 2. Advanced command
// Using user defined message  "-e EnvMessage=moimoi"
// and using user defined port number "-e EnvPort=8888"
$ docker run --rm -p 8888:8888 -e EnvPort=8888 -e EnvMessage=moimoi my_golang_echo_server:latest
```
