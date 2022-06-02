import { Media } from "./media";
import { Reaction } from "./reaction";

export class Post{
    postId: string;
	userId: string;
	creationTime:any;
	mediaAssets: Array<Media>
	postText:string;
	reactions:Array<Reaction>
	comments:Array<Comment>

    constructor(){

            this.postId='';
            this.userId='';
            this.creationTime=null;
            this.mediaAssets=[];
            this.postText='';
            this.reactions=[];
            this.comments=[];

    }

    
}