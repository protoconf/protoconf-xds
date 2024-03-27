FROM envoyproxy/envoy:dev-7f6df241c61094c6e1ec26c8d1acf05f5803a32c
COPY envoy.yaml /etc/envoy/envoy.yaml
RUN chmod go+r /etc/envoy/envoy.yaml
