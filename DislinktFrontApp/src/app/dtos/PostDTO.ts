export class PostDTO{
    postID:string;
    userEmail:string;
    creationTime:string;
    mediaAssets:Array<string>;
    postText:string;
    reactions:Array<string>;
    comments: Array<string>;

    constructor(_id:string,userEmail:string, creationtime:string,mediaassets:[],reactions:[], posttext:string, comments:[]){
        this.userEmail=userEmail;
        this.creationTime=creationtime;
        this.mediaAssets=mediaassets;
        this.comments = comments;
        this.postText = posttext;
        this.reactions =reactions;
        this.postID = _id
    }
}