export class UserDTO{
    name: string;
    surname: string;
    email: string;
    numtel: string;
    sex: string;
    bdday: number;
    bdmonth: number;
    bdyear: number;
    username:string;
    password:string;
    bio:string;
    exp: Array<string>;
    edu: Array<string>;
    interests: Array<string>;
    skills: Array<string>;
    isPrivate: boolean;
    constructor(name: string, surname: string, email:string, numtel:string,
        sex:string, bdday: number, bdmonth:number, bdyear:number, 
        username:string, password: string, bio:string, exp = [], 
        edu =[], intersets=[], skills=[], isPrivate = true) {
        this.name = name;
        this.surname = surname;
        this.email = email;
        this.numtel = numtel;
        this.bdday = bdday;
        this.bdmonth = bdmonth;
        this.bdyear = bdyear;
        this.sex = sex;
        this.username = username;
        this.password = password; 
        this.bio = bio;
        this.exp = exp;
        this.edu = edu;
        this.interests = intersets;
        this.skills = skills;
        this.isPrivate = isPrivate;
    }
}