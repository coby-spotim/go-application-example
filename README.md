# Example Go Application

This is an example Go application which can be used as a template to build Go projects.

## Running The Application

This project currently contains a single command and application, the `gin-api-example`.

In order to run the application, you can either use the Dockerfile and docker-compose file in the root of the repository, or run the (`cmd/gin-api-example/main.go`)[https://github.com/coby-spotim/go-application-example/blob/master/cmd/gin-api-example/main.go] file.

## Structure

The project follows parts of the Go Project structure set forth by the (Golang Project Layout)[https://github.com/golang-standards/project-layout].

The Golang Project Layout repository compiles and explains how many of the popular projects in the Go community structure their project, and sets forth a standard using these structures.

In the root of the project, you'll find a a few things to help you get started:
- A README file, which you are reading right now
- A .gitignore file, for obvious reasons
- A Gopkg.toml and a Gopkg.lock file, for dependency management
- A Dockerfile and a docker-compose.yml, for quick starting the application in a containerized state

You'll also find the following folders. See the explanations below for what each folder houses.

### `/cmd`

Here you'll find the code that will generate a binary to run. Code here should be small and concise, as all it needs to do is generate a binary for you to run. Do not put application code here, application code should live in either the `/internal` folder (private code not for use outside the current workspace) or the `/pkg` folder (public code available for reuse).

### `/configs`

Here you'll find the configuration files for the project. Currently there are simply some blank YAML files in there, but any sort of application configuration files can be added here.

### `/internal`

Here you'll find code that we do not want to export to the world. Any private application code that shouldn't need to be imported in any other project (routers, handlers, etc.)

### `/pkg`

Here you'll find code that can be reused in other projects (types, sdks, etc.)

### `/vendor`

The vendor folder which holds all of the dependencies that are being pulled via any of the Go dependency managers.