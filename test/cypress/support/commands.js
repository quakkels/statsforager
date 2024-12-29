// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add('login', (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add('drag', { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add('dismiss', { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite('visit', (originalFn, url, options) => { ... })

Cypress.Commands.add('login', email => {
	cy.visit('/login')
	cy.get('input[name="email"]').type(email)
	cy.get('button[type="submit"]').click()
	cy.get('h1').contains('Check email')
	cy.wait(5000)
	cy.readFile('cypress/fixtures/' + email + '.txt').then((content) => {
		const urlRegex = /(https?:\/\/[^\s]+)/;
		const match = content.match(urlRegex);
		const url = match ? match[0] : null;
		cy.visit(url)
			.url()
			.should("include","app/dashboard")
	});
	cy.wait(5000)
})
