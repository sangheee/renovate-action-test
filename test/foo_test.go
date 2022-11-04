package test

import (
	"fmt"
	"testing"
)

var testSts = `
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: kafka
spec:
  selector:
    matchLabels:
      app: kafka
  serviceName: kafka-broker
  replicas: 2
  template:
    metadata:
      labels:
        app: kafka
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: kubernetes.io/arch
                operator: In
                values:
                - amd64
                - arm64
      containers:
      - name: broker
        image: bitnami/kafka:3.0.2
        env:
        - name: HOSTNAME_VALUE
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        - name: KAFKA_LISTENERS_COMMAND
          value: "echo \"PLAINTEXT://${HOSTNAME}.kafka-broker.kafka.svc.cluster.local:9092\""
        - name: KAFKA_LOG4J_OPTS
          value: "-Dlog4j.configuration=file:/opt/kafka/config/log4j.properties"
        - name: BROKER_ID_COMMAND
          value: "echo ${HOSTNAME##*-}"
        - name: KAFKA_ZOOKEEPER_CONNECT
          value: "zk-cs.kafka.svc.cluster.local:2181"
        - name: KAFKA_LOG_DIRS
          value: "/var/lib/kafka/data/topics"
        - name: KAFKA_LOG_RETENTION_HOURS
          value: "-1"
        - name: KAFAK_ZOOKEEPER_CONNECTION_TIMEOUT_MS
          value: "6000"
        - name: KAFAK_AUTO_CREATE_TOPICS_ENABLE
          value: "false"
        ports:
        - containerPort: 9092
        livenessProbe:
          initialDelaySeconds: 60
          exec:
            command:
            - /bin/sh
            - -c
            - 'netstat -tulpn | grep LISTEN | grep 9092'
        volumeMounts:
        - name: data
          mountPath: /var/lib/kafka/data
  volumeClaimTemplates:
  - metadata:
      name: data
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 500Mi`

func Test(t *testing.T) {
	fmt.Println(testSts)
}
