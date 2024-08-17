curl --request PUT --url http://0.0.0.0:8000/api/sites/92dd13d0-8eff-4e02-9951-4f335602d99f/impressions/fe4e0125-74f5-4cc7-90e7-dae97d25c2e2 \
	--header 'content-type: application/json' \
	--data-binary @- << EOF
{
	"userAgent":"this agent",
	"language":"en",
	"location":"0.0.0.0:8000/SOMEWHERE",
	"referrer":null,
	"startedUtc":"2024-07-10T11:48:54.000Z",
	"completedUtc":"2024-07-10T11:48:54.000Z"
}
EOF

