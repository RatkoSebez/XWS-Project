export class NewPostDTO{
    links:Array<string>;
    photos:Array<string>;
    postText:string;

    constructor(links:[], photos:[], postText:string){
        this.links=links;
        this.photos=photos;
        this.postText=postText;
    }
}