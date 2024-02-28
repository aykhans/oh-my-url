# Run for development

## Rename and fill environment files

- Rename the following files:
    - `/app/config/.env.example` to `/app/config/.env`
    - `/config/postgresql/.env.example` to `/config/postgresql/.env`

## Run with PostgreSQL

- Run the following command to start the application with PostgreSQL:
    ```bash
    docker compose -f docker-compose-postgres.yml up
    ```