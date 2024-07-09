-- first sql migration file
-- Command to run files:
--	psql -h {host} -U {username} -d stats -a -f deploy.sql

BEGIN;
\i scripts/0000.sql
\i scripts/0001.sql
\i scripts/0002.sql
\i scripts/0003_impressions_add_is_leaving_pkey.sql
\i scripts/0004_impressions_site_fk.sql
\i scripts/0005_impressions_change_leaving.sql
COMMIT;

