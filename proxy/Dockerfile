FROM envoyproxy/envoy:v1.28-latest 
# distroless-v1.28-latest
WORKDIR /envoy
COPY envoy.yaml .
ENTRYPOINT [ "envoy", "-c", "envoy.yaml" ]
