import TodoList from '../TodoList.vue'

type TaskType = {
  id: number
  title: string
  completed: boolean
}

describe('TodoList', () => {
  const tasks: TaskType[] = [
    {
      id: 1,
      title: 'todo1',
      completed: false
    },
    {
      id: 2,
      title: 'todo2',
      completed: true
    },
    {
      id: 3,
      title: 'todo3',
      completed: false
    }
  ]

  // mount the component with the tasks prop
  beforeEach(() => cy.mount(TodoList, { props: { tasks } }))

  it('should render the input field and the add button', () => {
    cy.get('.custom-input')
      .should('be.visible')
      .should('have.attr', 'placeholder', 'Add a new todo')
    cy.get('.btn-primary').should('be.visible').should('have.text', 'Add')
  })

  it('should render the list of tasks', () => {
    cy.get('ul').should('be.visible')
    cy.get('ul').children().should('have.length', 3)
    cy.get('ul')
      .children()
      .each((child, index) => {
        cy.wrap(child).find('label').should('contain', tasks[index].title)
        cy.wrap(child)
          .find('input[type="checkbox"]')
          .should(tasks[index].completed ? 'be.checked' : 'not.be.checked')
      })
  })

  it('should add a new task when typing a title and clicking the add button', () => {
    cy.get('.custom-input').type('Write a test{enter}')
    cy.get('.btn-primary').click()
    cy.get('ul').children().should('have.length', 4)
    cy.get('ul')
      .children()
      .last()
      .find('label')
      .should('contain', 'Write a test')
    cy.get('ul')
      .children()
      .last()
      .find('input[type="checkbox"]')
      .should('not.be.checked')
  })

  it('should delete a task when clicking the delete button', () => {
    cy.get('ul').children().first().find('button').click()
    cy.get('ul').children().should('have.length', 2)
    cy.get('ul')
      .children()
      .first()
      .find('label')
      .should('contain', 'Do laundry')
    cy.get('ul')
      .children()
      .first()
      .find('input[type="checkbox"]')
      .should('be.checked')
  })

  it('should show a message when there are no tasks', () => {
    cy.get('ul')
      .children()
      .each((child) => {
        cy.wrap(child).find('button').click()
      })
    cy.get('ul').children().should('have.length', 0)
    cy.get('h2').should('be.visible').should('have.text', 'No Todos Here😞')
  })
})
