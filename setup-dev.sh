#!/bin/bash

docker run -d --restart=always --name pgsql -e POSTGRES_USER=cachet -e POSTGRES_PASSWORD=cachet postgres
docker run -d --restart=always --name cachet --link pgsql:pgsql -p 80:8000 -e DB_DRIVER=pgsql \
    -e DB_HOST=pgsql -e DB_DATABASE=cachet -e DB_USERNAME=cachet -e DB_PASSWORD=cachet \
    -e MAIL_DRIVER=smtp -e MAIL_HOST=smtp.gmail.com -e MAIL_PORT=465 -e MAIL_ENCRYPTION=ssl \
    -e APP_KEY=base64:5azW+xGOYjEX9Bq9RKksKvJlvNLbUXrUT4e0TpduS1g= \
    cachethq/docker:2.3.14
