export class JwtData{
    constructor(
        public email: string,
        public username: string,
        public token: string,
        public isAdmin: boolean,
        public agentToken:string
    ){}
}