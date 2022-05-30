import { SalaryForPosition } from "./dto/SalaryForPosition";
import { Job } from "./Job";

export class Company{
    constructor(
        public name: string,
        public address: string,
        public email: string,
        public phoneNumber: string,
        public description: string,
        public isApproved: boolean,
        public ownerEmail: string,
        public comments: string[],
        public interviewReviews: string[],
        public salaryForPosition: SalaryForPosition[],
        public jobs: Job[],
    ){}
}