# SRE-Bootcamp

This is a REST API server for CRUD operations on Students. The repo is used as the main service in One2N's [SRE Bootcamp](https://playbook.one2n.in/sre-bootcamp/sre-bootcamp-exercises)

### Local Setup

Use `make` to view all the available commands

```bash
❯ make
Usage: make <target>
Targets:
  help                           Display help
  dc-up                          dc-up starts the dependencies in the background
  dc-down                        dc-down stops the running dependencies
  build                          build builds the server binary
  run                            run starts the server
  test                           test runs the tests
  check-lint                     check-lint checks whether golangci-lint is installed
  lint                           lint runs the linter
  build-migration                build-migration builds the migration binary
  migrate                        migrate runs the up migration
  migrate-down                   migrate-down runs the down migration. You can optionally pass the number of steps to rollback like: make migrate-down steps=1
  force-migrate                  force-migrate force migrates a schema version. It requires a version to be passed like: make force-migrate version=1
```

- Clone the repo.
  
- `make dc-up` sets up the local dependencies. Make sure docker is installed on your system

- Create a `.env` file in the root of this repo and copy all the keys from `.env.default`. Add suitable values for your environment

- Once the dependencies are up and running, run the database migration using `make migrate`. Check `make help` for more info on force migration and rollback

- Use `make run` to run the API server.

- Import the `Student API.postman_collection.json` file in Postman to use the APIs


