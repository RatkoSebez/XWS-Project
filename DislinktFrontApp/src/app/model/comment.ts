import { ThisReceiver } from "@angular/compiler";

export class Comment{
    commentId:number;
	postId:number;
	userId:number;
	creationTime:any;
	comment:string;
    
    constructor(commentId:number,
        postId:number,
        userId:number,
        creationTime:any,
        comment:string){
            this.commentId=commentId;
            this.postId=postId;
            this.userId=userId;
            this.creationTime=creationTime;
            this.comment=comment;
    }
}