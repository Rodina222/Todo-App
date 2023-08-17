import TodoItem from '../TodoItem.vue'

type TaskType = {
  id: number
  title: string
  completed: boolean
}

describe('TodoItem', () => {
  const task: TaskType = {
    id: 1,
    title: 'Buy groceries',
    completed: false
  }

  beforeEach(() => cy.mount(TodoItem, { props: { task } }))

  it('should render the task title and checkbox', () => {
    cy.get('label').should('contain', task.title)
    cy.get('input[type="checkbox"]').should('not.be.checked')
  })

  it('should toggle the completed state when clicking the checkbox', () => {
    cy.get('input[type="checkbox"]').click()
    cy.get('input[type="checkbox"]').should('be.checked')
    cy.get('span').should('have.class', 'completed')
  })

  it('should edit the task title when double clicking the label', () => {
    cy.get('label').dblclick()
    cy.get('input[type="text"]').should('be.visible')
    cy.get('input[type="text"]').clear().type('Do homework{enter}')
    cy.get('label').should('contain', 'Do laundry')
  })

  it('should delete the task when clicking the delete button', () => {
    const deleteSpy = cy.spy().as('deleteSpy')

    cy.mount(TodoItem, { props: { task, onDelete: deleteSpy } })
    cy.get('button').click()
    cy.get('@deleteSpy').should('have.been.calledWith', task.id)
  })
})
