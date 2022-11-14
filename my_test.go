package main


var testPod2 = `
apiVersion: v1
kind: Pod
metadata:
  name: mysql
spec:
  containers:
  - name: mysql
    image: mysql:5.7.39
    ports:
    - containerPort: 80`
