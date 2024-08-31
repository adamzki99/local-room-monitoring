# Use the official Grafana image as the base image
FROM grafana/grafana:11.2.0

# Set environment variables for Grafana admin user and password
# Use default values to avoid issues if environment variables are not set
ENV GF_SECURITY_ADMIN_USER=${GF_SECURITY_ADMIN_USER:-admin}
ENV GF_SECURITY_ADMIN_PASSWORD=${GF_SECURITY_ADMIN_PASSWORD:-admin}

# Copy configuration files into the appropriate directories
COPY datasource.yaml /etc/grafana/provisioning/datasources/datasource.yaml
COPY dashboard.json /var/lib/grafana/dashboards/dashboard.json
COPY custom-dashboard.yaml /etc/grafana/provisioning/dashboards/custom-dashboard.yaml

# Expose the Grafana port
EXPOSE 3000

# Define the default command to run Grafana
CMD [ "grafana-server", "--homepath", "/usr/share/grafana", "cfg", "/etc/grafana/grafana.ini" ]
