function setupStatsForager(setup) {
	let impressionId = null;
	let statsForager = {

		arrive: function() {
			this.setGuid();
			let impression = {
				impressionId: impressionId,
				userAgent: navigator.userAgent,
				language: navigator.language,
				location: window.location.href,
				referrer: document.referrer
			}
			console.log(impression);
		},

		leave: async function() {
			console.log(`Impression ${impressionId} is ending.`);
			alert("interruptiong");
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
	};

	console.log("here");
	console.log(setup);
	addEventListener("beforeunload", async () => { await statsForager.leave(); });
	statsForager.arrive(setup);
};


