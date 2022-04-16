SELECT 'CREATE DATABASE er'
WHERE NOT EXISTS(SELECT FROM pg_database WHERE datname = 'er')
\gexec


DROP TABLE IF EXISTS currencies CASCADE;
CREATE TABLE currencies (
	id serial NOT NULL PRIMARY KEY,
	info json NOT NULL
);