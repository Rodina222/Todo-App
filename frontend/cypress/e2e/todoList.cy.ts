describe('TodoList', () => {
  it('allows a user to add and delete todos', () => {
    cy.visit('http://localhost:8080/')

    cy.get('.form-control').type('Buy milk').type('{enter}')

    cy.get('.todoItem')
      .should('have.length', 1)
      .first()
      .should('contain', 'Buy milk')

    cy.get('.form-control').type('Do laundry').type('{enter}')

    cy.get('.todoItem')
      .should('have.length', 2)
      .first()
      .should('contain', 'Buy milk')
      .next()
      .should('contain', 'Do laundry')

    cy.get('.todoItem button').first().click()

    cy.get('.todoItem')
      .should('have.length', 1)
      .first()
      .should('contain', 'Do laundry')
  })
})
