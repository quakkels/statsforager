import '../support/commands'

describe('Login', () => {

	it ('Unauthenticated user cannot access Dashboard', () => {
		cy.visit('app/dashboard')
			.url()
			.should("not.contain", "app/dashboard")
	});

	it('Authenticated user can access dashboard', () => {
		cy.login("me@example.com")
		cy.visit("app/dashboard")
			.url()
			.should("contain", "app/dashboard")
	})
})
