export class CommentDTO{
    postId: string;
    commentContent : string;

    constructor(id:string, content:string){
        this.postId=id;
        this.commentContent=content;
      
    }
}