export class CreateJobDTO{
    constructor(
        public companyName: string,
        public name: string,
        public description: string,
        public requirements: string,
        public benefits: string,
    ){}
}