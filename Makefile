help:
	@echo "    setup"
	@echo "        Build the development containers and install dependencies."
	@echo "    update"
	@echo "        Install / update dependencies in the containers."
	@echo "    bootstrap"
	@echo "        Set up and seed databases."
	@echo "        **IMPORTANT**: this wipes any existing data in your local database."
	@echo "    build"
	@echo "        Build the production images."

setup: dev-build update howto

dev-build:
	@docker-compose build

howto:
	@echo "Successfully built containers and installed dependencies."
	@echo "If this is your initial setup, you can run 'make bootstrap' next"
	@echo "to create and seed the database."

bootstrap:
	@echo "Bootstrapping Server service ..."
	@docker-compose run server bootstrap
	@echo ""
	@echo "You can now log into the development backend using the following credentials:"
	@echo ""
	@echo "Email: develop@offen.dev"
	@echo "Password: develop"
	@echo ""

update:
	@echo "Installing / updating dependencies ..."
	@docker-compose run homepage pip install --user -r requirements.txt
	@echo "Applying pending database migrations ..."
	@docker-compose run server migrate

DOCKER_IMAGE_TAG ?= latest
ROBOTS_FILE ?= robots.txt.staging

build:
	@docker build --build-arg siteurl=${SITEURL} --build-arg robots=${ROBOTS_FILE} -t offen/proxy:${DOCKER_IMAGE_TAG} -f build/proxy/Dockerfile .

.PHONY: setup build bootstrap build secret
