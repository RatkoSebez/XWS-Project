export class RegisterUserDto{
    constructor(
        public name: string,
        public surname: string,
        public email: string,
        public username: string,
        public password: string
    ){}
}