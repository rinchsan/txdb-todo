DB_DATABASE := db
DB_USER := root
DB_PASSWORD := password

.PHONY: up
up:
	cd etc/docker; docker-compose up -d

.PHONY: upbuild
upbuild:
	cd etc/docker; docker-compose up -d --build

.PHONY: down
down:
	cd etc/docker; docker-compose down

.PHONY: test
test:
	cd etc/docker; docker-compose exec --env DB_DATABASE=db_test api go test ./... -v

.PHONY: mysql
mysql:
	cd etc/docker; docker-compose exec mysql mysql -u${DB_USER} -p${DB_PASSWORD} ${DB_DATABASE}

.PHONY: logs
logs:
	cd etc/docker; docker-compose logs -f

.PHONY: gen
gen:
	XDG_CONFIG_HOME=etc sqlboiler mysql
	go get github.com/golang/mock/mockgen@v1.4.4
	go generate ./...
