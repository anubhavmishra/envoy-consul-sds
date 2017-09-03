FROM alpine:latest
ADD build/linux/amd64/envoy-consul-sds /envoy-consul-sds
CMD ["/envoy-consul-sds"]