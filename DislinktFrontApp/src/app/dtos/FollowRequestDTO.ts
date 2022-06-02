
export class FollowRequestDTO{
    SenderEmail: string;
    email : string;

    constructor(SenderEmail:string, ReceiverEmail:string){
        this.SenderEmail=SenderEmail;
        this.email=ReceiverEmail;
      
    }
}



