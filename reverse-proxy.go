package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"net/http"
	"net/http/httputil"
	"net/url"
	"gopkg.in/yaml.v2"
)

type DockerCompose struct {
	Services map[string]Service
}

type Service struct {
	ContainerName string            `yaml:"container_name"`
	Ports         []string
	Environment   []string
}

func ParseVariables(filename string) DockerCompose {
        // Read the docker-compose.yml file
        yamlFile, err := ioutil.ReadFile(filename)
        if err != nil {
                log.Fatalf("Failed to read file: %v", err)
        }

        // Parse the yaml file
        dockerCompose := DockerCompose{}
        err = yaml.Unmarshal(yamlFile, &dockerCompose)
        if err != nil {
                log.Fatalf("Failed to parse YAML: %v", err)
        }

	return dockerCompose
}


func MakeBackendMap(param_compose DockerCompose) map[string]*url.URL{

    backendURLs := make(map[string]*url.URL)

    for key, value := range param_compose.Services {
	if key == "reverse-proxy" {
	    continue
	}
	editedKey := "/"+key
	//fmt.Println(value.ContainerName)
        env_variable := strings.Split(value.Environment[0], "=")
	//fmt.Println("print split string: ", env_variable[0],env_variable[1])
	backendAppURL := fmt.Sprintf("http://%s:%s", key, env_variable[1])
	parsedURL, err := url.Parse(backendAppURL)
        if err != nil {
                panic(err)
        }

	backendURLs[editedKey]= parsedURL
    }

    return backendURLs
}

func main() {
        // Parsing docker-compose.yml file to info(container name, port, env variables)
	dockerCompose := ParseVariables("docker-compose.yml")
	backendURLs := MakeBackendMap(dockerCompose)
        for k,v := range backendURLs{
		fmt.Println("Possible path: "+k)
		fmt.Printf("Full URL: %s\n",v)
	}


	reverseProxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
		if backendURL, ok := backendURLs[req.URL.Path]; ok {
			req.URL.Scheme = backendURL.Scheme
			req.URL.Host = backendURL.Host}
		}}	

	listening_port := ":8080"	
	fmt.Println("Listening on ", listening_port)
	http.ListenAndServe(listening_port, reverseProxy)
}

