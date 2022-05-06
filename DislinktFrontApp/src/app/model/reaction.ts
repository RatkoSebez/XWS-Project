export class Reaction{
    reactionId:number;
	postId:number;
	userId:number;
	creationTime:any;
	reactionType:string;

    constructor(reactionId:number,
        postId:number,
        userId:number,
        creationTime:any,
        reactionType:string){
            this.reactionId=reactionId;
            this.postId=postId;
            this.userId=userId;
            this.creationTime=creationTime;
            this.reactionType=reactionType;
        }
}