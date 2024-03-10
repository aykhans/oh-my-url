#!/bin/bash

if [ -z "$CASSANDRA_USERNAME" ] || [ -z "$CASSANDRA_PASSWORD" ]; then
    echo "Error: Username or password environment variables are not set."
    exit 1
fi

until cqlsh -u cassandra -p cassandra -e "DESCRIBE KEYSPACES" ohmyurl-cassandra-1 || cqlsh -u $CASSANDRA_USERNAME -p $CASSANDRA_PASSWORD -e "DESCRIBE KEYSPACES" ohmyurl-cassandra-1; do
    echo >&2 "Cassandra is unavailable - sleeping"
    sleep 1
done

echo >&2 "Cassandra is up - executing command"

cqlsh -u cassandra -p cassandra -e "CREATE ROLE IF NOT EXISTS $CASSANDRA_USERNAME WITH PASSWORD = '$CASSANDRA_PASSWORD' AND SUPERUSER = true AND LOGIN = true;" ohmyurl-cassandra-1
cqlsh -u $CASSANDRA_USERNAME -p $CASSANDRA_PASSWORD -e "ALTER ROLE cassandra WITH PASSWORD='abcdef' AND SUPERUSER=false AND LOGIN = false;" ohmyurl-cassandra-1
cqlsh -u $CASSANDRA_USERNAME -p $CASSANDRA_PASSWORD -e "CREATE KEYSPACE IF NOT EXISTS $CASSANDRA_KEYSPACE WITH REPLICATION = {'class': 'NetworkTopologyStrategy', 'datacenter1': 3};" ohmyurl-cassandra-1

exit 0
