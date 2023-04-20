#!/bin/bash

# Define here a bash script that will be executing the migration on the database(es).
# Don't hardcode any credentials here. Use the credentials which are supplied through 
# TaskDefinition (see environment and secrets list from ./deployment/task-definition.json.tpl)

# This script should always exit with 0 if migration succeeded, and with non-zero if it fails
# defaultSchema is created by flyway implicitly before first migration script execution
env | grep GORM_EXPERIMENT
flyway migrate \
  -defaultSchema=${GORM_EXPERIMENT_DB_SCHEMA} \
  -user=${GORM_EXPERIMENT_DB_USERNAME} \
  -password=${GORM_EXPERIMENT_DB_PASSWORD} \
  -url=jdbc:postgresql://${GORM_EXPERIMENT_DB_HOST}:${GORM_EXPERIMENT_DB_PORT}/${GORM_EXPERIMENT_DB_NAME}?${GORM_EXPERIMENT_DB_PARAMS} \
  -connectRetries=5 \
  -locations='filesystem:/migrations'

