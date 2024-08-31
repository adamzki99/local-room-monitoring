# Use the latest official PostgreSQL image
FROM postgres:15

# Set environment variables for PostgreSQL
# Use default values to avoid issues if environment variables are not set
ENV POSTGRES_USER=${POSTGRES_USER:-postgres}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD:-password}
ENV POSTGRES_DB=${POSTGRES_DB:-mydatabase}

# Copy initialization SQL scripts to the Docker image
# PostgreSQL will automatically run these scripts on container start
COPY init.sql /docker-entrypoint-initdb.d/
COPY data.sql /docker-entrypoint-initdb.d/

# Expose PostgreSQL default port
EXPOSE 5432