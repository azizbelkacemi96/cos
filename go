replicaCount: 3

image:
  repository: my-image
  tag: latest
  pullSecret: my-image-pull-secret

serviceAccount:
  name: my-service-account

containers:
  - name: my-container
    ports:
      - containerPort: 9898
    command: ["/bin/my-command"]
    args: ["arg1", "arg2"]
    volumeMounts:
      - name: config-volume
        mountPath: /etc/config
        readOnly: true
      - name: certs-volume
        mountPath: /etc/certs
        readOnly: true
    resources:
      limits:
        cpu: "1"
        memory: "512Mi"
      requests:
        cpu: "0.5"
        memory: "256Mi"

volumes:
  - name: config-volume
    secret:
      secretName: my-config-secret
  - name: certs-volume
    secret:
      secretName: my-certs-secret
