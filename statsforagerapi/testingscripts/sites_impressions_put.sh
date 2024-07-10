curl --request PUT --url http://0.0.0.0:8000/api/sites/mysite/impression/thisimpression \
	--header 'content-type: application/json' \
	--data-binary @- << EOF
{
	"impressionId":"thisimpression",
	"userAgent":"this agent",
	"language":"en",
	"location":"0.0.0.0:8000/SOMEWHERE",
	"referrer":null,
	"startedUtc":null,
	"completedUtc":null
}
EOF

