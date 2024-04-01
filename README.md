# SRE-Bootcamp

This is a REST API server for CRUD operations on Students. The repo is used as the main service in One2N's [SRE Bootcamp](https://playbook.one2n.in/sre-bootcamp/sre-bootcamp-exercises)

### Requirements

- Docker - https://docs.docker.com/get-docker
- Docker Compose - https://docs.docker.com/compose/install
- Make - https://askubuntu.com/questions/161104/how-do-i-install-make

### Local Setup

Use `make` to view all the available commands

```zsh
❯ make
Usage: make <target>
Targets:
  help                           Display help
  dc-up                          dc-up starts the dependencies in the background
  dc-down                        dc-down stops the running dependencies
  build                          build builds the server binary
  build-app                      build-app builds the app docker image
  run                            run starts the server
  run-app                        run-app starts the app container on port 9090
  test                           test runs the tests
  check-lint                     check-lint checks whether golangci-lint is installed
  lint                           lint runs the linter
  build-migration                build-migration builds the migration binary
  migrate                        migrate runs the up migration
  migrate-down                   migrate-down runs the down migration. You can optionally pass the number of steps to rollback like: make migrate-down steps=1
  force-migrate                  force-migrate force migrates a schema version. It requires a version to be passed like: make force-migrate version=1
  docker-build-migration         docker-build-migration builds the migration docker image
  docker-migrate                 docker-migrate runs the migration docker container
  docker-migrate-down            docker-migrate-down runs the down migration in the migration container. You can optionally pass the number of steps to rollback like: make docker-migrate-down steps=1
  local-setup                    local-setup sets up the local environment in docker
  local-teardown                 local-teardown tears down the local environment in docker
```

- Clone the repo.

- Create a `.env` file in the root of this repo and copy all the keys from `.env.default`. Add suitable values for your environment.

- Use `make local-setup` to run the app along with its dependencies in docker.

- Import the `Student API.postman_collection.json` file in Postman to use the APIs.

### Running Migration

- Create a `migration.env` file in the root of this repo. Add the following variables:
  - DSN=<YOUR DSN>
  - ENVIRONMENT=dev
  - MIGRATION_PATH=db/migrations
  - *Note : You might need to use `host.docker.internal` as your hostname in the DSN for certain environments (e.g: macOS, windows)*

- Use `make docker-migrate` and `make docker-migrate-down` for up and down migrations. You can check for more details on these commands by hitting `make`.