import { Media } from "./media";
import { Reaction } from "./reaction";

export class Post{
    postId: number;
	userId: number;
	creationTime:any;
	mediaAssets: Array<Media>
	postText:string;
	reactions:Array<Reaction>
	comments:Array<Comment>

    constructor(){

            this.postId=-1;
            this.userId=-1;
            this.creationTime=null;
            this.mediaAssets=[];
            this.postText='';
            this.reactions=[];
            this.comments=[];

    }

    
}