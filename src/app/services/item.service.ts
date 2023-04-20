import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Item } from '../item';
import { map } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class ItemService {
    
    getItem(itemID : number) : Item {
        let singleItem : Item = new Item(NaN, "", "", NaN, "");
        console.log("Attempting to get an item.");
        this.http.get<any>('/item').pipe(
            map(data => ({
                ID : data.ID,
                name : data.name,
                description : data.description,
                uidOfOwner : data.uidOfOwner,
                imgSrc : data.imgSrc
            })
            )).subscribe(object =>
                {
                    singleItem = object;
                }
            );
        return singleItem;
    }

    setDefaultValues(item : Item) : void {
        item.ID = -1;
        item.uidOfOwner = -1;
        item.name = "default";
        item.description = "placeholder";
        item.imgSrc = "../../assets/shiba2.jpg";
    }

  constructor(private http : HttpClient) {}
}
