import HTTPMethod from "http-method-enum";
import Item from "../models/item";
import ItemAdapter from "../models/itemAdapter";
import Service from "./service";
import { Remote } from './common/common'

/**
 * Provides functions to call the server only for completed items.
 * There are some missing functions implementations because this
 * is out of the scope of this project.
 */
export default class CompletedItemService implements Service<Item> {
    get(id: number): Promise<Item> {
        throw new Error("Method not implemented.");
    }

    async getAll(): Promise<Item[]> {
        const response = await fetch(`${Remote.Common.getBaseUrl()}/items/completed`, { method: HTTPMethod.GET })

        if (response.status === 200) {
            return (await response.json()).map((item: ItemAdapter) => Item.create(item))
        } else {
            throw new Error('Could not retrieve the list of completed todos')
        }
    }

    create(item: Item): Promise<Item> {
        throw new Error("Method not implemented.");
    }

    update(item: Item): Promise<Item> {
        throw new Error("Method not implemented.");
    }

    delete(id: number): Promise<number> {
        throw new Error("Method not implemented.");
    }
}