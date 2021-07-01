FROM ubuntu:20.10

RUN apt-get update \
    && apt-get install -y nginx \
		&& rm -rf /var/lib/apt/lists/*

# Add the OTel C++ nginx module for instrumentation
# GitHub: open-telemetry/opentelemetry-cpp-contrib - instrumentation/nginx
# Pulled from an artifact of a merge commit build: https://github.com/open-telemetry/opentelemetry-cpp-contrib/actions/runs/949633629
COPY ./otel_ngx_module.so /usr/lib/nginx/modules
# Maybe in the future these artifacts will be on project GitHub releases like:
# ADD https://github.com/open-telemetry/opentelemetry-cpp-contrib/releases/download/1.0.0/otel_ngx_module-ubuntu-20.10-stable.so /usr/lib/nginx/modules

# Configs
COPY ./otel-conf.toml /etc/otel-conf.toml
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./proxies.conf /etc/nginx/sites-enabled/default

CMD ["nginx", "-g", "daemon off;"]
