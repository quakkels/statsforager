describe('Signup', () => {
  it('Navigate to signup from login', () => {
    cy.visit('/login')
		cy.get('main a[href="/register"]').click()
		cy.url().should('include', '/register')
  })
})
