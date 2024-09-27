# StatsForager To Do

## Mission

1. Make a website analytics service that respects site owner's data, user privacy, and doesn't depend on Big Tech.
2. Promote an independent internet by offering alternative to Big Tech services.

### In Progress
- dashboard site select
- create site

### Needs doing
- locations with top ten growth
- site impression count over time
- browsers visiting your site
- location visits
- top referrers and top referred destination
- API connect to database from docker container
- marketing: "Traffic Analysis" copy for landing page
- marketing: blog section of statsforager
- marketing: "Actionable Insights" copy for landing page
- Write "Getting Started" (probably on landing page)

### Done
- move statsforager client side script to statsforagerweb to be servered from there rather than from the client's server
- add charts to dashboard (apexcharts?)
- registration & login rate limiting (tollbooth?)
- csrf (nosurf)
- logout and logout header
- middleware to verify authenticated user
- registration + login
- user session (scs)
- send email function
- dashboard page + template
- landing page route + template
- forager site template (main template with support for partials, and place to put custom template functions)
- serve static files for client css and scripts
- validate impression location is approved for site
- hosting deploy server config allowing CORS
- script client call api
- api (save sitekey, impressionId, IP, user data from client)
- update database to track length of impression
- hosting deploy process (deploy api to railway via docker hub)
- create and migrate remote database
- choose host provider
- script client create session id
- script client collect user data (impressionId, PageLoadDateTime, PageExitDateTime, Referrer, URL, user-agent)
- choose database

## Beyond MVP
- page keyword analysis on most popular pages (user story: As a site owner, I want to understand what keywords are working the best for my pages, so that I can make informed decisions on new content for my site.)
- site ownership verification
- api noscript gif using site key
- api endpoint for script client using site key (is this necessary? maybe not MVP)
- donate button
- subscribe
- how to proxy guide
- getting started documentation
- script client GDPR support toggle
- script client GDPR opt out
- noscript client (gif)
