#!/bin/bash

# Exit script on any error
set -e

# Load environment variables from the .env file
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo ".env file not found!"
  exit 1
fi

# Define the file location for datasource.yaml 
file_location="$(pwd)/grafana/datasource.yaml"

# Remove old file if it exists
if [ -f "$file_location" ]; then
  rm "$file_location"
fi

# Create a new datasource.yaml file
touch "$file_location"

# Prepare the output string
output_string=$(cat <<EOF
apiVersion: 1

datasources:
  - name: lrm-postgres
    type: postgres
    url: lrm-postgres:5432
    user: $DATABASE_GRAFANA_USER
    secureJsonData:
      password: '$DATABASE_GRAFANA_PASSWORD'
    jsonData:
      database: $POSTGRES_DB
      sslmode: 'disable'
      maxOpenConns: 100
      maxIdleConns: 100
      maxIdleConnsAuto: true
      connMaxLifetime: 14400
      postgresVersion: 1500
      timescaledb: false
EOF
)

# Write the output string to the file
echo "$output_string" > "$file_location"

# Print a message indicating success
echo "Created file datasource.yaml in directory $(pwd)/grafana/"
