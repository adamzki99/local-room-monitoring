[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=adamzki99_local-room-monitoring&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=adamzki99_local-room-monitoring)

# local-room-monitoring
Small project using a database and some micro-controllers to store and display the living factors of rooms.

## TODO

### Base

#### data-collector

- [ ] Add API endpoint to add a new device to the database
  - Make it so the address should be updated if the device_id is already present in the table. DHCP should not make the device go offline.

#### database

- [x] Change so that the database is built with a script so that all schemas are set automatically
- [ ] Change the database writes to make them safe from SQL injections

### edge

- [ ] Add edge device implementation

### DevOps

- [ ] Setup more Github actions:
  - [x] SonarCloud
  - [ ] CodeSene
  - [ ] etc.

### General
- [x] Add Graphana
- [ ] HTTPS/TLS support?
- [ ] Improve code test coverage
- [X] Add documentation
  - [ ] Architecture
  - [ ] Database tables
  - [ ] etc.
- [ ] Cleanup...
- [ ] Add Prometheus (might not happen)
