# StatsForager To Do

## MVP
**Goal:** Make a webiste analytics service that respects site owner privacy, user privacy,
and removes dependencies on unstrustworthy data hogs like Google.

### In Progress
- hosting deploy process (deploy api to railway via docker hub)
- api (save sitekey, impressionId, IP, user data from client)

### Needs doing
- API connect to database from docker container
- script client call api
- api noscript gif using site key
- api endpoint for script client using site key
- user dashboard
- hosting deploy server config allowing CORS
- site ownership verification

### Done
- create and migrate remote database
- choose host provider (railway)
- script client create session id
- script client collect user data (impressionId, PageLoadDateTime, PageExitDateTime, Referrer, URL, user-agent)
- choose database

## Beyond MVP

- donate button
- subscribe
- how to proxy guide
- getting started documentation
- script client GDPR support toggle
- script client GDPR opt out
- noscript client (gif)
