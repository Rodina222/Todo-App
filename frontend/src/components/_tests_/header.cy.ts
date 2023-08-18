import HelloToDo from "../Header.vue";

describe("Page Header", () => {
  beforeEach(() => cy.mount(HelloToDo));

  it("should have header title", () => {
    cy.get(".header h1").should("exist");
  });

  it("should have header subtitle", () => {
    cy.get(".header h3").should("exist");
  });
});

