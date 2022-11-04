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

func Test(t *testing.T) {
	fmt.Println(testPod)
}
