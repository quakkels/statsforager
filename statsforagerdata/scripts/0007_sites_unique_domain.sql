DO $$
DECLARE
	-- start upgrade to this version
	next_version INTEGER := 7;
BEGIN
	
	IF EXISTS(SELECT * FROM db_version WHERE version >= next_version)
	THEN
		-- we're on or beyond this version already
		RAISE NOTICE 'Skipping version %', next_version;
	ELSE
		-- execute migration
		ALTER TABLE sites
			ADD CONSTRAINT unique_owner_account_domain UNIQUE (owner_account, domain);

		-- we're done. set current version
		UPDATE db_version SET version = next_version;
	END IF;
END $$;

