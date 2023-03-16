package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

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


//func mustParseURL(rawURL string) *url.URL {
//        parsedURL, err := url.Parse(rawURL)
//        if err != nil {
//                panic(err)
//        }
//        return parsedURL
//}


func main() {
	// Read the docker-compose.yml file
	yamlFile, err := ioutil.ReadFile("docker-compose.yml")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Parse the yaml file
	dockerCompose := DockerCompose{}
	err = yaml.Unmarshal(yamlFile, &dockerCompose)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	for k,v := range dockerCompose.Services {
	        fmt.Println("Print services")
		fmt.Println(k, v)
	}

	path, err := os.Getwd()
        if err != nil {
	    log.Println(err)
        }
	fmt.Println("")
	fmt.Println("mypath:",path)  
//	backendApp1URL := fmt.Sprintf("http://%s:8765", os.Getenv("BACKEND_APP1_NAME"))
//	backendApp2URL := fmt.Sprintf("http://%s:8000", os.Getenv("BACKEND_APP2_NAME"))
//
//	backendURLs := map[string]*url.URL{
//		"/app1": mustParseURL(backendApp1URL),
//		"/app2": mustParseURL(backendApp2URL),
//	}
//
//	reverseProxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
//		if backendURL, ok := backendURLs[req.URL.Path]; ok {
//			req.URL.Scheme = backendURL.Scheme
//			req.URL.Host = backendURL.Host
//		}
//	}}
//
//	http.ListenAndServe(":8080", reverseProxy)
}

