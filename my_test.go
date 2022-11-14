package main


var testPod2 = `
apiVersion: v1
kind: Pod
metadata:
  name: mysql
spec:
  containers:
  - name: mysql
    image: mysql:8.0.31
    ports:
    - containerPort: 80`
