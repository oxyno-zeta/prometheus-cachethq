# Developer

## General

This golang project is based on Go Modules and mustn't be placed in GOPATH to avoid linter problems with imports.

This project is mainly a GraphQL server but everything can be used to build a REST API or a cronjob.

The `static/` folder is here to allow static web hosting.

All configurations are all in the `conf/` folder. All configurations can be put in different files. They will be watched for changes and the application will reload them. All services in the application can be notified of configuration change.

The `graphql/` folder contains all graphql files that contains the GraphQL Schema. These will be parsed and managed by [gqlgen](https://gqlgen.com/).

The project is linted by [GolangCI-Lint](https://golangci-lint.run/) and built by [Goreleaser](https://goreleaser.com/) or [Gox](https://github.com/mitchellh/gox).

The project have a convention in the coding strategy. All the business code related to your application specifically is located in a specific folder. They will be called business unit.

## Convention

- the git commit convention is the angular one (see [here](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#-commit-message-guidelines))
- Editorconfig is used to keep file content in a uniform way

## Install

This project is using the python software called `pre-commit`. This is used to install and have git pre-commit hooks.

Those ones are here to validate code, lint projects, lint and validate GraphQL, ...

Moreover, some tools are used in the backend project. These tools are using NodeJS and Yarn for package installation.

Just run the script called `./install-deps.sh` in order to install needed dependencies.

## Release

In order to release, the project is using `semantic-release` in order to generate a release directly with the Changelog (git based).

## Available features

- Prometheus metrics available on `/metrics`
- Health checks available on `/health`
- Ready endpoint available on `/ready`
  - This one will check if health checks are valid by default and only when a SIGTERM or a SIGINT is caught, the endpoint will be marked as Service Unavailable
- The application will caught SIGTERM and SIGINT and will stop the application when no primary requests are in progress

## Structure

The `cmd` folder contains all main packages.

The `pkg` folder contains all packages used by the application.

In this folder, there is:

- `pkg/../authx`: This folder contains packages related to authentication and authorization check.
- `pkg/../business`: This folder contains all business units of your application. These will contain services, models and data access object (dao) methods.
- `pkg/../common`: This folder contains common errors and utils used in all other packages.
- `pkg/../config`: This folder contains the package managing configuration. This provide a manager that give access to the last configuration loaded in the application. This allow to add hook for configuration reload.
- `pkg/../database`: This folder contains the package managing the SQL database connection and access.
- `pkg/../lockdistributor`: This contains a package that allow to acquire a distributed semaphore based on PostgreSQL.
- `pkg/../log`: This contains a package to have a logger.
- `pkg/../metrics`: This contains a package for metrics (Prometheus in this case).
- `pkg/../server`: This package contains servers code, GraphQL code and utils.
- `pkg/../tracing`: This package allow to have trace in the application using OpenTracing (implementation done with Jaeger).
- `pkg/../version`: This contains a package to have built version of the current application.
