# StatsForager To Do

## MVP
**Goal:** Make a webiste analytics service that respects site owner privacy, user privacy,
and removes dependencies on unstrustworthy data hogs like Google.

### In Progress
- validate impression location is approved for site

### Needs doing
- api noscript gif using site key
- API connect to database from docker container
- user dashboard
- site ownership verification

### Done
- hosting deploy server config allowing CORS
- script client call api
- api (save sitekey, impressionId, IP, user data from client)
- update database to track length of impression
- hosting deploy process (deploy api to railway via docker hub)
- create and migrate remote database
- choose host provider (railway)
- script client create session id
- script client collect user data (impressionId, PageLoadDateTime, PageExitDateTime, Referrer, URL, user-agent)
- choose database

## Beyond MVP
- api endpoint for script client using site key (is this necessary? maybe not MVP)
- donate button
- subscribe
- how to proxy guide
- getting started documentation
- script client GDPR support toggle
- script client GDPR opt out
- noscript client (gif)
