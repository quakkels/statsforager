const { defineConfig } = require("cypress");

module.exports = defineConfig({
	e2e: {
		specPattern: '**/*.cy.js',
		setupNodeEvents(on, config) {
			// implement node event listeners here
		},
		baseUrl: 'http://localhost:8000/'
	},
});
