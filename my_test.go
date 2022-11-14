package main

var somePackageImageVersion = "v0.2.0" // update: depName=your.registry.address/pkgName

var someK8sDeployTemplate= `
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: app
  name: app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: app
  template:
    metadata:
      labels:
        app: app
    spec:
      containers:
        - image: your.registry.address/pkgName:tag`

var someConstruct = &struct{
      version  string
    }{
			version:  "v0.2.0", // update: datasource=github-release depName=foo/bar
		}
