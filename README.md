# Health Information System

## Technologies used

- Golang 1.24 <https://go.dev/dl/>
- Make <https://www.gnu.org/software/make/>
- PostgreSQL <https://www.pgadmin.org/download/>
- Docker <https://docs.docker.com/desktop/>
- Air (live reload) <https://github.com/air-verse/air>

## Powerpoint presentation:

## Database design:

## API testing: 

## Live API: 

### Setup local development

1. Clone the repository

   ```bash
   git clone https://github.com/wathika-eng/hims --depth 1
   # cd into the directory
   cd hims
   ```

2. Install dependencies

   ```bash
   make deps
   # or with go
   go mod tidy
   ```

3. Create a `.env` file

   ```bash
    cp .env.example .env
    ```

4. Create a PostgreSQL database

    ```bash
    # create a database with the name in .env
    createdb -U postgres -h localhost -p 5432 hims
    ```

5. Run the application with live reload

    ```bash
    air
    ```

6. With docker:

    ```bash
    COMPOSE_BAKE=true docker-compose up -d
    ```
    <!-- or
    ```bash
    docker build -t hims .
    docker run -p 8080:8080 hims
    docker run --env-file .env -v $(pwd)/.env:/root/.env hims:latest
    ``` -->

#### Resources used

<https://12factor.net/>
<https://bun.uptrace.dev/guide/golang-orm.html>
<https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/>
