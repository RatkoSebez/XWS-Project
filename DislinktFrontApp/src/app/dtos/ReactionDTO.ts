export class ReactionDTO{
    react:string;
    userEmail:string;
    postID:string


    constructor( type:string, usermail:string, id:string){
        this.react=type;
        this.userEmail=usermail;
        this.postID = id ;
      
    }
}