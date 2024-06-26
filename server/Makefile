# 移行が終わるまでのつなぎ
migrate_kari:
	PGPASSWORD=${DB_PASSWORD} psqldef -h db -p ${DBPORT} -U ${DB_USER} ${DB_NAME} < _tools/postgres/migration/create_table.sql

migrate:
	PGPASSWORD=${DB_PASSWORD} psqldef -h db -p ${DBPORT} -U ${DB_USER} ${DB_NAME} < _tools/postgres/init/0000_init.sql
