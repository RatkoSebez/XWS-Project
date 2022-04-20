export class AuthResponseDTO{
    token: string;
    email: string;
    roles: Array<string>;

    constructor(token: string, email: string, roles =[]) {
        this.token = token;
        this.email = email;
        this.roles = roles;
    }
}