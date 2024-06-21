-- upgrade to version 2

DO $$
BEGIN
	IF EXISTS(SELECT * FROM db_version WHERE version >= 1)
	THEN
		-- we're on or beyond this version already
		RAISE NOTICE 'Skipping version 1';
	ELSE
		-- execute migration
		ALTER TABLE sites DROP CONSTRAINT sites_pkey;
		ALTER TABLE sites ADD PRIMARY KEY (site_key);
		ALTER TABLE sites ADD site_name VARCHAR(255);

		CREATE TABLE impressions (
			impression_id UUID NOT NULL,
			ip_address INET NOT NULL,
			user_agent VARCHAR(510) NOT NULL,
			language VARCHAR(24),
			location VARCHAR(2000) NOT NULL,
			referrer VARCHAR(2000),
			date_time timestamp default (timezone('utc', now())) NOT NULL,
			PRIMARY KEY(impression_id)
		);

		-- we're done. set current version
		UPDATE db_version SET version = 1;
	END IF;
END $$;
