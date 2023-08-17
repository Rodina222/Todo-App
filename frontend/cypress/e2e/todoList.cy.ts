describe('TodoList', () => {
  beforeEach(() => {
    cy.request({
      method: 'POST',
      url: 'localhost:8096/todos',
      body: {
        name: 'todo'
      }
    })

    cy.visit('localhost:8080')
  })

  it('should add a new task', () => {
    const taskTitle = 'New Task'

    cy.get('.custom-input').type(taskTitle)
    cy.get('.btn-primary').click()

    cy.get('.todoItem').should('contain', taskTitle)
  })

  it('should delete a task', () => {
    const taskIndex = 0

    cy.get('.todoItem').eq(taskIndex).as('taskItem')
    cy.get('@taskItem')
      .invoke('text')
      .then((taskText) => {
        cy.get('@taskItem').find('button').click()
        cy.get('.todoItem').should('not.contain', taskText)
      })
  })
  it('should delete a task', () => {
    const taskIndex = 0

    cy.get('.todoItem').eq(taskIndex).as('taskItem')
    cy.get('@taskItem')
      .invoke('text')
      .then((taskText) => {
        cy.get('@taskItem').find('button').click()
        cy.get('.todoItem').should('not.contain', taskText)
      })
  })

  it('should mark a task as completed', () => {
    const taskIndex = 0

    cy.get('.todoItem').eq(taskIndex).as('taskItem')
    cy.get('@taskItem').find('.todoCheckbox').check()

    cy.get('@taskItem')
      .find('.todoText')
      .should('have.css', 'text-decoration', 'line-through solid rgb(0, 0, 0)')
  })

  it('should update a task', () => {
    const taskIndex = 0
    const updatedTaskTitle = 'Updated Task'

    cy.get('.todoItem').eq(taskIndex).as('taskItem')
    cy.get('@taskItem').find('.todoText').dblclick()
    cy.get('@taskItem')
      .find('.todoText.editing')
      .clear()
      .type(updatedTaskTitle)
      .type('{enter}')

    cy.get('@taskItem').should('contain', updatedTaskTitle)
  })
})
