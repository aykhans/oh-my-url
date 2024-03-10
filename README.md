## Run with PostgreSQL

- Rename and fill environment files:
    - `/app/config/postgres.env.example` to `/app/config/postgres.env`
    - `/config/postgresql/.env.example` to `/config/postgresql/.env`

- Run the following command to start the application with PostgreSQL:
    ```bash
    docker compose -f docker-compose-postgres.yml up
    ```

## Run with Cassandra

- Rename and fill environment files:
    - `/app/config/cassandra.env.example` to `/app/config/cassandra.env`
    - `/config/cassandra/init-cassandra.env.example` to `/config/cassandra/init-cassandra.env`

- Run the following command to start the application with PostgreSQL:
    ```bash
    docker compose -f docker-compose-cassandra.yml up
    ```