#passwords #generative-files #environment-variables

This documentation provides a summary of how files in the project utilize environment variables for configuration and employ Bash scripts to generate resources necessary for database and Grafana datasource creation.

## 1. Configuration via Environment Variables

### 1.1 Usage of Environment Variables

Environment variables, specified in a `.env` file in the _./base/_ directory, are referenced within configuration files to inject values.

Use the following snippet as a template for the ```.env``` file.

```bash
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=

DATABASE_SCHEMA=
DATABASE_GRAFANA_USER=
DATABASE_GRAFANA_PASSWORD=

GF_SECURITY_ADMIN_USER=
GF_SECURITY_ADMIN_PASSWORD=
```

## 2. Resource Generation with Bash Scripts

### 2.1 Database Creation

The creation of the database makes use of a ```init.sql``` file for the basic configuration. Run the following command to create it.

```bash
$ ./create_init_sql.sh
```

### 2.2 Grafana Datasource Setup

The creation of the Grafana datasource makes use of a ```datasource.yaml``` file for the basic configuration. Run the following command to create it.

```bash
$ ./create_datasource_yaml.sh
```
