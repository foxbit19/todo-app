import HTTPMethod from "http-method-enum";
import Item from "../models/item";
import Service from "./service";
import ItemAdapter from '../models/itemAdapter';

/**
 * ItemService implements a Service using Item model as
 * type argument
 */
export default class ItemService implements Service<Item> {
    /**
     * Gets the base url
     * @returns a string of the base url to use in this service
     */
    private getBaseUrl() {
        return process.env.REACT_APP_BASE_URL;
    }

    async get(id: number): Promise<Item> {
        const response = await fetch(`${this.getBaseUrl()}/items/${id}`, { method: HTTPMethod.GET })

        if (response.status === 200) {
            return Item.create(await response.json())
        } else {
            throw new Error(`Could not retrieve this todo: ${id}`)
        }
    }

    async getAll(): Promise<Item[]> {
        const response = await fetch(`${this.getBaseUrl()}/items/`, { method: HTTPMethod.GET })

        if (response.status === 200) {
            return (await response.json()).map((item: ItemAdapter) => Item.create(item))
        } else {
            throw new Error('Could not retrieve the list of todos')
        }
    }

    async create(item: Item): Promise<Item> {
        const response = await fetch(`${this.getBaseUrl()}/items/`, { method: HTTPMethod.POST, body: JSON.stringify(item) })

        if (response.status === 200) {
            return item
        } else {
            throw new Error('Could not create this item')
        }
    }

    async update(item: Item): Promise<Item> {
        const response = await fetch(`${this.getBaseUrl()}/items/`, { method: HTTPMethod.PUT, body: JSON.stringify(item) })

        if (response.status === 200) {
            return item
        } else {
            throw new Error('Could not update this item')
        }
    }

    async delete(id: number): Promise<number> {
        const response = await fetch(`${this.getBaseUrl()}/items/${id}`, { method: HTTPMethod.DELETE })

        if (response.status === 200) {
            return id
        } else {
            throw new Error(`Could not delete this item: ${id}`)
        }
    }
}