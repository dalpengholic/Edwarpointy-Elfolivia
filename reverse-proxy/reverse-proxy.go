package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"net/http"
	"net/http/httputil"
	"net/url"
//	"os"

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
//	for k,v := range dockerCompose.Services {
//	        fmt.Println("Print services")
//		fmt.Println(k, v)
//	}

//        fmt.Println("")
//	fmt.Println("print dockerCompose.Services: ", dockerCompose.Services)
//        fmt.Println("")
//        fmt.Println("print dockerCompose.Services[app1]: ", dockerCompose.Services["app1"])	
//        fmt.Println("")
//        fmt.Println("print dockerCompose.Services[app1].ContainerName: ", dockerCompose.Services["app1"].ContainerName)
//        fmt.Println("")
//        fmt.Println("print dockerCompose.Services[app1].Ports: ", dockerCompose.Services["app1"].Ports)
//        fmt.Println("")
//        fmt.Println("print dockerCompose.Services[app1].Environment: ", dockerCompose.Services["app1"].Environment)
//        fmt.Println("")
//        fmt.Println("print dockerCompose.Services[app1].Environment[0]: ", dockerCompose.Services["app1"].Environment[0])
//        fmt.Println("")
//	env_variable := strings.Split(dockerCompose.Services["app1"].Environment[0], "=")
//	fmt.Println("print split string: dockerCompose.Services[app1].Environment[0]: ", env_variable)
//	fmt.Println("print select first string: ", env_variable[0])
        fmt.Println("++++++++++++++++++++++++++")
	backendURLs := MakeBackendMap(dockerCompose)
        for k,v := range backendURLs{
	    fmt.Println(k)
	    fmt.Println(v)
	}

	//path, err := os.Getwd()
        //if err != nil {
	//    log.Println(err)
        //}
	//fmt.Println("")
	//fmt.Println("mypath:",path)  
//	backendApp1URL := fmt.Sprintf("http://%s:8765", os.Getenv("BACKEND_APP1_NAME"))
//	backendApp2URL := fmt.Sprintf("http://%s:8000", os.Getenv("BACKEND_APP2_NAME"))
//
//	backendURLs := map[string]*url.URL{
//		"/app1": mustParseURL(backendApp1URL),
//		"/app2": mustParseURL(backendApp2URL),
//	}
//
	reverseProxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
		if backendURL, ok := backendURLs[req.URL.Path]; ok {
			req.URL.Scheme = backendURL.Scheme
			req.URL.Host = backendURL.Host
		}
	}}

	http.ListenAndServe(":8080", reverseProxy)
}

