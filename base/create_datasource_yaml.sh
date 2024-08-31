#!/bin/bash

# Load the environment variables from the .env file
export $(grep -v '^#' .env | xargs)

# Set the file location for datasource.yaml 
fileLocation="$(pwd)/grafana/datasource.yaml"

# Remove file if old instance already exists
rm $fileLocation

# Create the datasource.yaml file
touch $fileLocation

output_string="apiVersion: 1

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
      timescaledb: false"

echo "$output_string" >> "$fileLocation"
echo "" >> "$fileLocation" 


# Print a message indicating success
echo "Created file datasource.yaml in directory $(pwd)/database/"
