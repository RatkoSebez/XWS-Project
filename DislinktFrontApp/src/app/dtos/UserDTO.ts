export class UserDTO{
    name: string;
    surname: string;
    email: string;
    numtel: string;
    sex: string;
    bdateday: number;
    bdatemonth: number;
    bdateyear: number;
    username:string;
    password:string;
    bio:string;
    experience: Array<string>;
    education: Array<string>;
    interests: Array<string>;
    skills: Array<string>;
    isprivate: boolean;
    constructor(name: string, surname: string, email:string, numtel:string,
        sex:string, bdday: number, bdmonth:number, bdyear:number, 
        username:string, password: string, bio:string, exp = [], 
        edu =[], intersets=[], skills=[], isprivate = true) {
        this.name = name;
        this.surname = surname;
        this.email = email;
        this.numtel = numtel;
        this.sex = sex;

        this.bdateday = bdday;
        this.bdatemonth = bdmonth;
        this.bdateyear = bdyear;
        this.username = username;
        this.password = password; 
        this.bio = bio;
        this.experience = exp;
        this.education = edu;
        this.interests = intersets;
        this.skills = skills;
        this.isprivate = isprivate;
    }
}