package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var host, url string
	flag.StringVar(&host, "host", "hz.gl", "HTTP/S host name")
	flag.StringVar(&url, "url", "/", "URL to HTTP GET")
	var port int
	flag.IntVar(&port, "port", 80, "HTTP/S port number")
	var useTLS bool
	flag.BoolVar(&useTLS, "tls", false, "Use TLS without verification on the specified port")
	flag.Parse()

	if host == "" || port < 1 {
		panic("Please specify both host and port")
	}

	client := &http.Client{Timeout: 30 * time.Second}

	if useTLS {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	total := 10
	endpoint := fmt.Sprintf("%s:%d%s", host, port, url)
	if useTLS {
		endpoint = "https://" + endpoint
	} else {
		endpoint = "http://" + endpoint
	}
	for i := 0; i < total; i++ {
		fmt.Printf("(%d of %d) Sending GET request to %s using a timeout of 30 seconds\n", i+1, total, endpoint)
		time.Sleep(1 * time.Second)

		resp, err := client.Get(endpoint)
		if err == nil {
			fmt.Printf("HTTP client got response code %d, first 200 bytes of content:\n", resp.StatusCode)
			body, bodyErr := ioutil.ReadAll(resp.Body)
			if bodyErr == nil {
				if len(body) > 200 {
					body = body[:200]
				}
				fmt.Println(string(body))
			} else {
				fmt.Printf("Failed to read response content: %+v\n", bodyErr)
			}
			resp.Body.Close()

		} else {
			fmt.Printf("HTTP client error: %+v\n", err)
		}
		fmt.Println()
	}
}
