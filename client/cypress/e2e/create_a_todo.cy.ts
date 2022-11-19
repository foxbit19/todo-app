describe('Creating a message', () => {
    it('Displays the message in the list', () => {
        cy.visit('http://localhost:3000');

        cy.get('[data-testid="todo_description"]')
            .type('My first todo');

        cy.get('[data-testid="todo_save"]')
            .click();

        cy.get('[data-testid="todo_description"]')
            .should('have.value', '');

        cy.contains('My first todo');
    });
});