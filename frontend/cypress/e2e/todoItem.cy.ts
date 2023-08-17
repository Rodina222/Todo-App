describe('TodoItem', () => {
  beforeEach(() => {
    cy.visit('localhost:8096/todos')
  })
  it('renders the component correctly', () => {
    cy.get('li').should('exist')
    cy.get('input[type="checkbox"]').should('exist')
    cy.get('span').should('exist')
    cy.get('button').should('exist')
  })

  it('toggles the editing mode when double-clicking the label', () => {
    cy.get('label').dblclick()
    cy.get('input[type="text"]').should('exist')

    cy.get('label').dblclick()
    cy.get('input[type="text"]').should('not.exist')
  })

  it('updates the task when pressing Enter in editing mode', () => {
    const newTitle = 'Updated Task'

    cy.get('label').dblclick()
    cy.get('input[type="text"]').type(newTitle).type('{enter}')
    cy.get('span').should('contain', newTitle)
  })

  it('updates the task when checkbox is clicked', () => {
    cy.get('input[type="checkbox"]').click()
    cy.get('span').should('have.class', 'completed')

    cy.get('input[type="checkbox"]').click()
    cy.get('span').should('not.have.class', 'completed')
  })

  it('emits "delete" event when delete button is clicked', () => {
    cy.get('button').click()
    // Assert the "delete" event is emitted
    cy.get('@deleteEvent').should('have.been.calledOnce')
    // Assert the correct task ID is passed in the event
    cy.get('@deleteEvent').should('have.been.calledWith', 1) // Replace 1 with the actual task ID
  })

  // Add more test cases for other functionality if needed
})
