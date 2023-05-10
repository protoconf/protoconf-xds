FROM envoyproxy/envoy:dev-334a50e7ab10ff3247c300fdf32a2425e9deaa55
COPY envoy.yaml /etc/envoy/envoy.yaml
RUN chmod go+r /etc/envoy/envoy.yaml
