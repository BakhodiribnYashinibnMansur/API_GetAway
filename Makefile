CURRENT_DIR=$(shell pwd)

permission-sh:
	sudo chmod +x ./scripts/gen-proto.sh

proto-gen:
	./scripts/gen-proto.sh	${CURRENT_DIR}

clone-protos:
	rm -rf protos/* && cp -R ur_protos/* protos

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge

migrate-up:
	docker run --mount type=bind,source="${CURRENT_DIR}/migrations,target=/migrations" --network ${NETWORK_NAME} migrate/migrate \
		-path=/migrations/ -database=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable up

migrate-down:
	docker run --mount type=bind,source="${CURRENT_DIR}/migrations,target=/migrations" --network ${NETWORK_NAME} migrate/migrate \
		-path=/migrations/ -database=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable down

swag-init:
	swag init -g api/main.go -o api/docs
.PHONY: proto
