function setupStatsForager(setup) {
	let impressionId = null;
	let impression = null;
	let statsForager = {

		arrive: function() {
			this.setGuid();
			let dateUtc = this.getDateUtcNowIso();
			impression = {
				impressionId: impressionId,
				userAgent: navigator.userAgent,
				language: navigator.language,
				location: window.location.href,
				referrer: document.referrer,
				startedUtc: dateUtc,
				completedUtc: dateUtc
			}
			window.fetch(
				`http://localhost:8000/api/sites/${setup.siteKey}/impressions/${impressionId}`,
				{
					method: 'PUT',
					headers: {
						"Content-Type": "application/json",
					},
					body: JSON.stringify(impression)
				}
			);
		},

		leave: function() {
			impression.completedUtc = this.getDateUtcNowIso();
			window.fetch(
				`http://localhost:8000/api/sites/${setup.siteKey}/impressions/${impressionId}`,
				{
					method: 'PUT',
					headers: {
						"Content-Type": "application/json",
					},
					body: JSON.stringify(impression)
				}
			);
		},

		setGuid: function() {
			if (crypto.randomUUID) {
				impressionId = crypto.randomUUID();
				return;
			}

			let rand = crypto.getRandomValues(new Uint16Array(8));
			let i = 0;
			impressionId = "00-0-4-1-000".replace(/[^-]/g,
				s => (rand[i++] + s * 0x10000 >> s).toString(16).padStart(4, "0")
			);
		},

		getDateUtcNowIso: function() {
			let now = new Date();
			let nowUtc = Date.UTC(
				now.getUTCFullYear(), 
				now.getUTCMonth(),
				now.getUTCDate(), 
				now.getUTCHours(),
				now.getUTCMinutes(), 
				now.getUTCSeconds());
			let nowUtcFmt = new Date(nowUtc).toISOString();
			return nowUtcFmt;
		},
	};

	addEventListener("beforeunload", async () => { statsForager.leave(); });
	statsForager.arrive(setup);
};


