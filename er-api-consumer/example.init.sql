SELECT 'CREATE DATABASE er'
WHERE NOT EXISTS(SELECT FROM pg_database WHERE datname = 'er')
\gexec

DROP TABLE IF EXISTS rates CASCADE;
CREATE TABLE rates (
	id serial PRIMARY KEY,
	result VARCHAR ( 50 ) UNIQUE NOT NULL,
	documentation VARCHAR ( 50 ) NOT NULL,
	terms_ofu_se VARCHAR ( 255 ) UNIQUE NOT NULL,
	time_last_update_unix TIMESTAMP NOT NULL,
	time_last_update_utc VARCHAR ( 50 ) NOT NULL,
    time_next_update_unix TIMESTAMP NOT NULL,
	time_next_update_utc VARCHAR ( 50 ) NOT NULL,
	base_code VARCHAR ( 50 ) NOT NULL,
	conversion_rates TEXT []
);
