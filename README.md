# IQM-assignment
## IQM-assignment-3
### Main Concept
- 1. Making a reverse proxy server to have a single HTTP endpoint
  - To have a simple reverse proxy server, I avoided to bring `Traefik` or `Nginx` well-known reverse proxy
  - Echo Servers have to be set and ready behind the reverese proxy so that a alient only communicates with the reverse proxy server 

- 2. Parsed `docker-compose.yml` used during the build of reverse proxy
  - If you add more service blocks at docker-compose.yml file and edit `.env`, the reverse-proxy could parse the yaml file and can route to the new instances after rebuilding 

- 3. `.env` file used for convenience 

### How to use
- 1. Open terminal and create a folder and download this branch in the folder:
```Shell
$ mkdir mytmp
$ cd mytmp
$ git clone git@github.com:dalpengholic/IQM-assignment.git
$ cd IQM-assignment
```
- 2. Checkout to task-3 branch
`$ git checkout task-3`

- 3. Create `.env` file
`$ vim .env`
```Shell
WEBSERVER1_NAME=backapp1
WEBSERVER2_NAME=backapp2
```

- 4. Build docker image
`$ docker-compose up -d`

- 5. If you want to add more echo server
  - 1. Add more service block to docker-compose file
```
...
  app3:
    image: webserver-6
    container_name: ${WEBSERVER3_NAME}
    hostname: ${WEBSERVER3_NAME}
    environment:
      - EnvPort=8883
      - EnvMessage=hihihoho
    networks:
      - backend
```
  - 2. Add env variabe at .env
```
WEBSERVER1_NAME=backapp1
WEBSERVER2_NAME=backapp2
WEBSERVER3_NAME=backapp3
```
  - 3. Rerun with --build flag
`$ docker-compose up -d --build`

## IQM-assignment-2
### Main concept of Dockerfile
- Using multi-stage builds to create a Minimum-Viable Product (MVP) Image
  - 1st stage
     - To build binary file, copy only necessary files like `go.mod, webserver.go and READEME.md` from localhost
  - 2nd stage
     - Copy from the files from 1st stage to busybox image

- The size of mvp image decreased to `11.5MB` from 314MB of golang build image  

### How to build an echo server from Dockerfile?
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
$ docker run --rm -p 8000:8000 my_golang_echo_server:latest

// 2. Advanced command
// Using user defined message  "-e EnvMessage=moimoi"
// and using user defined port number "-e EnvPort=8888"
$ docker run --rm -p 8888:8888 -e EnvPort=8888 -e EnvMessage=moimoi my_golang_echo_server:latest
```

## IQM-assignment-1
Decision Point-1: Which langauge would be better to implement simple echo server?
- Candidates: Golang, Python
- After a bit of searching for each language, I ended up choosing Golang.
- Compared to Python, with Golang it seemed easier to implement an Echo server through the default libraries without any web framework. Also, the number of lines of code required to write seemed short. 
