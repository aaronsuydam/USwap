import { Injectable, platformCore } from "@angular/core";

export class Item {
    ID : number;
    uidOfOwner : number;
    
    name : string;
    description : string;
    imgSrc : string
    
    constructor(ID : number = NaN, name : string = "", description : string = "", uidOfOwner : number = NaN, imgSrc : string = "") {
        this.ID = ID;
        this.uidOfOwner = uidOfOwner;
        this.name = name;
        this.description = description;
        this.imgSrc = imgSrc;
    };

    

}