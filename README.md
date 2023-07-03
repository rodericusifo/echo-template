# ECHO TEMPLATE

Application which can ...

## RUN APPLICATION ON DOCKER

- Step 1: Install Docker. Docs [Link](https://docs.docker.com/get-docker/)

- Step 2: Copy `docker.application.env.example` to `docker.application.env` in `environments` folder
  
`INSTRUCTION` >> fill the empty values, from this:

```properties
# docker.application.env
...
# DATABASE SQL
...
DATABASE_POSTGRES_NAME=""
DATABASE_POSTGRES_USERNAME=""
DATABASE_POSTGRES_PASSWORD=""
...
```

to this:

```properties
# docker.application.env
...
# DATABASE SQL
...
DATABASE_POSTGRES_NAME="db_name_example"
DATABASE_POSTGRES_USERNAME="db_username_example"
DATABASE_POSTGRES_PASSWORD="db_password_example"
...
```

- Step 3: Run docker compose build

```properties
# OR `docker-compose --env-file environments/docker.application.env build`
docker compose --env-file environments/docker.application.env build
```

- Step 4: Run docker compose up

```properties
# OR `docker-compose --env-file environments/docker.application.env up -d`
docker compose --env-file environments/docker.application.env up -d
```

## RUN APPLICATION ON DEV ENVIRONMENT

Run this command

```properties
make start-dev
```

## DOCUMENTATION

See in folder `docs` **OR** run on Postman below

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/10344918-27d85a45-7c41-4b84-9a31-9af5ab1e7a87?action=collection%2Ffork&collection-url=entityId%3D10344918-27d85a45-7c41-4b84-9a31-9af5ab1e7a87%26entityType%3Dcollection%26workspaceId%3D667868fa-663b-45d5-a9ec-252ff52cb9c8#?env%5BGo%20Echo%20App%20Template%20Env%5D=W3sia2V5IjoiYmFzZV91cmwiLCJ2YWx1ZSI6ImxvY2FsaG9zdDo4MDgxIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQiLCJzZXNzaW9uVmFsdWUiOiJsb2NhbGhvc3Q6ODA4MSIsInNlc3Npb25JbmRleCI6MH0seyJrZXkiOiJiYXNlX3VybF9kb2NrZXIiLCJ2YWx1ZSI6ImxvY2FsaG9zdDo4MDgyIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQiLCJzZXNzaW9uVmFsdWUiOiJsb2NhbGhvc3Q6ODA4MiIsInNlc3Npb25JbmRleCI6MX0seyJrZXkiOiJhY2Nlc3NfdG9rZW4iLCJ2YWx1ZSI6IiIsImVuYWJsZWQiOnRydWUsInR5cGUiOiJkZWZhdWx0Iiwic2Vzc2lvblZhbHVlIjoiZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBaQ0k2SWpobFlUYzNPR0pqTFRNNU5UZ3ROR1U1WmkwNFptRXlMV0U0WVRsaFpEaG1NbUZpTVNJc0ltNWhiV1VpT2lKaFpHMXBiaUlzSW1WdFlXbHNJam9pWVdSdGFXNUEuLi4iLCJzZXNzaW9uSW5kZXgiOjJ9XQ==)

## TEST COVERAGE

- Step 1: Install Mockery

```properties
go install github.com/vektra/mockery/v2@v2.20.0
```

- Step 2: Run this command

```properties
make test-cover
```

## NOTE

- Number 1: make sure you do all step in installation guide
- Number 2: the server should be listening on localhost:`<PORT>` \
  `NOTE`: you can see the `<PORT>` in docker-compose.yml on golang_container's ports

```yaml
# docker-compose.yaml
version: "3.9"
services:
  app:
    container_name: golang_container
    tty: true
    build: .
    ports:
      # <PORT>:<CONTAINER-PORT>
      - 8081:8080
...
```
