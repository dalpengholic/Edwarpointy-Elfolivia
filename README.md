## IQM-assignment-3
### Main Concept
To create a single HTTP endpoint, I built a reverse proxy server. To keep things simple, I avoided using popular reverse proxies such as Traefik or Nginx. I made sure to set up and configure the Echo Servers behind the reverse proxy so that all external communication would go through the reverse proxy server.

To ensure that the reverse proxy server could route traffic to new instances, I made the reverse proxy instance could parse the docker-compose.yml file when it runs. Additionally, I used the .env file to manage the environment variables required for the setup and configuration of the reverse proxy.

### How to use
- 1. Open terminal and create a folder and download this branch in the folder:
```Shell
$ mkdir mytmp
$ cd mytmp
$ git clone git@github.com:dalpengholic/IQM-assignment.git
```
- 2. Checkout to task-3 branch
```Shell
$ cd IQM-assignment
$ git checkout task-3
```

- 3. Create `.env` file
`$ vim .env`
```Shell
WEBSERVER1_NAME=app1
WEBSERVER2_NAME=app2
```

- 4. Build docker image
```Shell
$ docker-compose up -d
```

- 5. Check with curl command
```Shell 
$ curl http://localhost:8080/app1
hello from app1
$ curl http://localhost:8080/app2
hello from app2
```

- 6. If you want to add more echo server
  - 1. Add more service block to docker-compose file
```
...
  newapp3:
    image: my_webserver
    container_name: ${WEBSERVER3_NAME}
    hostname: ${WEBSERVER3_NAME}
    environment:
      - EnvPort=8883
      - EnvMessage=moi
    networks:
      - backend
```
  - 2. Add env variabe at .env
```
WEBSERVER1_NAME=app1
WEBSERVER2_NAME=app2
WEBSERVER3_NAME=newapp3
```
  - 3. Rerun with --build flag
`$ docker-compose up -d --build`


  - 5. Check with curl command
```Shell 
$ curl http://localhost:8080/app1
hello from app1
$ curl http://localhost:8080/app2
hello from app2
$ curl http://localhost:8080/app3
moi from newapp3
```

