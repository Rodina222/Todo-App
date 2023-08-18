/// <reference types="cypress" />
describe('Test Header', () => {
  it('should display the header1 with "✍️ ToDo Application" text and header3 with "What do you plan on doing🙂?" text', () => {
    cy.visit('http://localhost:8080/')
    cy.get('.header h1').should('have.text', '✍️ ToDo Application')
    cy.get('.header h3').should('have.text', 'What do you plan on doing🙂?')
  })
})
