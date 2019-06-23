# offen
[![CircleCI](https://circleci.com/gh/offen/offen/tree/master.svg?style=svg)](https://circleci.com/gh/offen/offen/tree/master)
[![Pivotal Tracker](https://img.shields.io/static/v1.svg?label=Project+Planning&message=Pivotal+Tracker&color=informational)](https://www.pivotaltracker.com/n/projects/2334535)
[![Patreon](https://img.shields.io/static/v1.svg?label=patreon&message=donate&color=e85b46)](https://www.patreon.com/offen)

> The offen analytics software

This repository contains all source code needed to build and run __offen__, both on the server as well as on the client. See each of the READMEs in the subdirectories for instructions on how to work on that particular area of the application.

---

Development of __offen__ has just started, so instructions are rare and things will stay highly volatile for quite some while. Also __do not use the software in its current state__ as it is still missing crucial pieces in protecting the data end to end.

Guidelines for running and developing the Software will be added when it makes sense to do so.

Project planning and issue tracking is done using [Pivotal Tracker](https://www.pivotaltracker.com/n/projects/2334535), but feel free to open a GitHub issue if you have a question or found a bug.

### Developing the application

The development setup requires `docker` and `docker-compose` to be installed.

After cloning the repository, you can build the containers and install dependencies using:

```sh
$ make setup
```

Next, create a local encryption key for the `kms` service and seed the database for the `server` application:

```sh
$ make bootstrap
```

You can test your setup by starting the application:

```sh
$ docker-compose up
```

which should enable you to access <http://localhost:8080/status> seeing a successful response.

### License

MIT © [offen](https://www.offen.dev)
