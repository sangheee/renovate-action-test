package main

import (
	"fmt"
	"testing"
)

var testPod = `
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
  - name: nginx
    image: nginx:1.14.2
    ports:
    - containerPort: 80`

type Spec struct {
	Image string `json:"image"`
}

type NewSpec struct {
	Registry string `json:"registry"`
	Name     string `json:"name"`
	Version  string `json:"version"`
}

func Test(t *testing.T) {
	spec := &Spec{
		Image: `hello-world:nanoserver-1809`, // update: versioning=regex:^(?<compatibility>.*)-(?<major>\d+)$
	}
	nSpec := &NewSpec{
		Registry: "",
		Name:     "hello-world",
		Version:  "nanoserver-1803", // update: versioning=regex:^(?<compatibility>.*)-(?<major>\d+)$
	}
	nnSpec := &NewSpec{
		Registry: "",
		Name:     "bitnami/kafka",
		Version:  "3.0.2",
	}

	fmt.Println(testPod, spec, nSpec, nnSpec)
}
