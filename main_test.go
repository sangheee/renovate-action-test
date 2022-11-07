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

func Test(t *testing.T) {
	spec := &Spec{
		Image: `hello-world:nanoserver-1709`, // update: versioning=regex:^(?<compatibility>.*)-(?<major>\d+)$
	}
	fmt.Println(testPod, spec)
}
