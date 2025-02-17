FROM envoyproxy/envoy:v1.33.0

# Instalar dependencias para protoc
RUN apt-get update && apt-get install -y \
    protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

# Copiar el archivo .proto al contenedor
COPY service.proto /etc/envoy/service.proto

# Generar el archivo .pb
RUN protoc -I /etc/envoy --include_imports --include_source_info \
    --descriptor_set_out=/etc/envoy/proto.pb /etc/envoy/service.proto
