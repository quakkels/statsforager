
-- get top X most popular pages
select 
	count(location) as impression_count,
	location
from impressions
where 
	site_key = '92dd13d0-8eff-4e02-9951-4f335602d99f'
	and
	(timezone('utc', now()) - interval '1 month') < started_utc
group by location
order by impression_count desc
limit 10;


-- get popularity by week over the last year
SELECT 
    location,
    DATE_TRUNC('week', started_utc) AS week_start,
    COUNT(1) AS visit_count
FROM 
    impressions
WHERE 
    started_utc >= CURRENT_DATE - INTERVAL '1 year'
GROUP BY 
    location, DATE_TRUNC('week', started_utc)
ORDER BY 
    week_start;
