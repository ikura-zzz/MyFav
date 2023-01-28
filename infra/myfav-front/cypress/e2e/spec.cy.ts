/// <reference types="Cypress" />
describe("ログイン関連", () => {
  beforeEach(() => {
    cy.visit("/");
  });
  it("サインイン失敗", () => {
    cy.get("[data-testid='username'").type("testuser", { force: true });
    cy.get("[data-testid='userpasswd'").type("testpassword", { force: true });
    cy.get("[data-testid='signinbutton'").click();
    cy.get("[data-testid='errormsg']").should(
      "contain.text",
      "このユーザーは存在しません。"
    );
  });
  it("サインアップ失敗", () => {
    cy.get("[data-testid='signup']").click();
    cy.get("[data-testid='newusername']").type("testuser", { force: true });
    cy.get("[data-testid='newpassword']").type("testpassword", { force: true });
    cy.get("[data-testid='retypepassword']").type("wrongpassword", {
      force: true,
    });
    cy.get("[data-testid='signupbutton']").click();
    cy.get("[data-testid='errormsg']").should(
      "contain.text",
      "パスワードが一致しません。"
    );
  });
  it("サインアップ成功", () => {
    cy.get("[data-testid='signup']").click();
    cy.get("[data-testid='newusername']").type("testuser", { force: true });
    cy.get("[data-testid='newpassword']").type("testpassword", { force: true });
    cy.get("[data-testid='retypepassword']").type("testpassword", {
      force: true,
    });
    cy.get("[data-testid='signupbutton']").click();
    cy.get("[data-testid='errormsg']").should("not.exist");
  });
  it("サインイン成功", () => {
    cy.get("[data-testid='username'").type("testuser", { force: true });
    cy.get("[data-testid='userpasswd'").type("testpassword", { force: true });
    cy.get("[data-testid='signinbutton'").click();
  });
});
