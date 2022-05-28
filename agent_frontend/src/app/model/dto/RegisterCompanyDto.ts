export class RegisterCompanyDto{
    constructor(
        public name: string,
        public address: string,
        public email: string,
        public phoneNumber: string,
        public description: string,
        public ownerEmail: string
    ){}
}