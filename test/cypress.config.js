const { defineConfig } = require("cypress");

module.exports = defineConfig({
	e2e: {
		specPattern: '**/*.cy.js',
		setupNodeEvents(on, config) {
			// implement node event listeners here
			on('task', {
				log(args) {
					console.log(...args)
					return null;
				}
			});
		},
		baseUrl: 'http://localhost:8000/'
	},
});
