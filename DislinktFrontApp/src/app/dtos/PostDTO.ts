export class PostDTO{
    links:Array<string>;
    photos:Array<string>;
    postText:string;
    comments: Array<string>;

    constructor(links:[], photos:[], postText:string, comments:[]){
        this.links=links;
        this.photos=photos;
        this.postText=postText;
        this.comments = comments;
    }
}