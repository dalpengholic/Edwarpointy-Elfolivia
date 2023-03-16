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
    image: my_webserver
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

