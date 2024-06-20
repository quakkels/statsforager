CREATE TABLE IF NOT EXISTS db_version (
	version INT
);


DO $$
BEGIN
	IF EXISTS(SELECT * FROM db_version WHERE version >= 0)
	THEN
		-- we're on or beyond this version already
		RAISE NOTICE 'Skipping version 0';
	ELSE
		-- migrate to version 0
		CREATE TABLE accounts (
			email VARCHAR(255) NOT NULL,
			is_active BOOLEAN NOT NULL,
			PRIMARY KEY(email)
		);
	
		CREATE TABLE sites (
			domain VARCHAR(255) NOT NULL,
			site_key UUID NOT NULL,
			owner_account VARCHAR(255) NOT NULL,

			PRIMARY KEY(domain),
			CONSTRAINT fk_accounts
				FOREIGN KEY(owner_account)
					REFERENCES accounts(email)
		);

		-- we're done. set current version
		INSERT INTO db_version VALUES (0);
	END IF;
END $$;
