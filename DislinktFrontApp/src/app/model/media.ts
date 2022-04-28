export class Media{
    mediaID:number;
    site:string;
    filepath:string;
    
    constructor(mediaID:number,
        site:string,
        filepath:string){
            this.mediaID=mediaID;
            this.site=site;
            this.filepath=filepath;
        }
}