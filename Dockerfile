FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.18-openshift-4.11 AS builder
WORKDIR /go/src/github.com/cgwalters/kvm-device-plugin
COPY . .
RUN make

FROM registry.access.redhat.com/ubi8/ubi:latest
COPY --from=builder /go/src/github.com/cgwalters/kvm-device-plugin/cmd/kvm/kvm /usr/bin/device-plugin-kvm
CMD ["/usr/bin/device-plugin-kvm"]
