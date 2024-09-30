package test

import (
	"fmt"
	"github.com/hwUltra/fb-tools/curl"
	"testing"
)

func TestGet(t *testing.T) {
	client := curl.NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}

	dataBytes, err := client.Get("/iaas/appio/ab", data)
	fmt.Println("Get: ", string(dataBytes), err)

	dataBytes, err = client.Curl("/iaas/appio/ab", curl.Get, data, curl.JsonType)
	fmt.Println("curl-get: ", string(dataBytes), err)
}

func TestPostJson(t *testing.T) {
	client := curl.NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}

	dataBytes, err := client.PostByForm("/iaas/appio/ab", data)
	fmt.Println("Post: ", string(dataBytes), err)

	dataBytes, err = client.Curl("/iaas/appio/ab", curl.Post, data, curl.FormType)
	fmt.Println("curl-post: ", string(dataBytes), err)
}

func TestPutJson(t *testing.T) {
	client := curl.NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}

	dataBytes, err := client.Put("/iaas/appio/ab", data)
	fmt.Println("Put: ", string(dataBytes), err)

	dataBytes, err = client.Curl("/iaas/appio/ab", curl.Put, data, curl.JsonType)
	fmt.Println("curl-put: ", string(dataBytes), err)
}

func TestPatchJson(t *testing.T) {
	client := curl.NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}

	dataBytes, err := client.Patch("/iaas/appio/ab", data)
	fmt.Println("Patch: ", string(dataBytes), err)

	dataBytes, err = client.Curl("/iaas/appio/ab", curl.Patch, data, curl.JsonType)
	fmt.Println("curl-patch: ", string(dataBytes), err)
}

func TestDelete(t *testing.T) {
	client := curl.NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}

	dataBytes, err := client.Delete("/iaas/appio", data)
	fmt.Println("Delete: ", string(dataBytes), err)

	dataBytes, err = client.Curl("/iaas/appio", curl.Delete, data, curl.JsonType)
	fmt.Println("curl-delete: ", string(dataBytes), err)
}
