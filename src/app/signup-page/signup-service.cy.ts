import { User } from '../interfaces/UserInterface';
import { SignupService } from './signup.service';


describe("SignupService.addUser()", () => {
    let signupService: SignupService;
    it("should add a user and return it", () => {
        const newUser: User = {
            username: "erob",
            email: "evan.robinson@ufl.edu",
            password: "test"
        };
        cy.request('POST', 'test', newUser).then(
            (res) => {
                expect(res.status).to.equal(201)
                expect(res.body).to.equal(newUser)
            }
        )
    });
});