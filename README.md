# Fork of original https://github.com/kubevirt/kubernetes-device-plugins

This repository is an updated fork of the original
https://github.com/kubevirt/kubernetes-device-plugins
which is now maintained as part of the larger [KubeVirt](https://github.com/kubevirt) project.

However, for some use cases one just wants "/dev/kvm in pods" and not the full
KubeVirt system.  Specifically at least for CoreOS, our [coreos-assembler](https://github.com/coreos/coreos-assembler/)
project is "kubernetes ready" and runs in e.g. OpenShift without special
privileges, as long as `/dev/kvm` is mounted.

You can use this project on bare metal systems, AWS `m5.metal` instances,
or [GCP](https://cloud.google.com/compute/docs/instances/enable-nested-virtualization-vm-instances).

# Usage

At the current time we are not publishing a pre-built image, so
you will need to build one and replace the `image` below.

Create a YAML file with the following, then use `kubectl create`.

```
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    name: device-plugin-kvm
  name: device-plugin-kvm
spec:
  selector:
    matchLabels:
      name: device-plugin-kvm
  template:
    metadata:
      labels:
        name: device-plugin-kvm
    spec:
      containers:
      - name: device-plugin-kvm
        image: <replace with built image>
        securityContext:
          privileged: true
        volumeMounts:
          - name: device-plugin
            mountPath: /var/lib/kubelet/device-plugins
      volumes:
        - name: device-plugin
          hostPath:
            path: /var/lib/kubelet/device-plugins
```

## Filtering nodes

Note that this daemonset must run with `privileged: true`, but that's
true of all device plugins.  The code here is very targeted.  However,
in a multi-tenant environment you may have concerns about applying
this daemonset everywhere.  It's possible to use e.g. a
[node selector](https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes/)
to only deploy the daemonset on a subset of nodes which require KVM
if you want to further isolate these workloads.
