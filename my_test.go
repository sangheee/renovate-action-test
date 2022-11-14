package main


var testPod2 = `
apiVersion: v1
kind: Pod
metadata:
  name: etcd
spec:
  containers:
  - name: etcd
    image: bitnami/etcd:3.5.5
    ports:
    - containerPort: 80`
