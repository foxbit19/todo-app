import HTTPMethod from "http-method-enum";
import Item from "../models/item";
import Service from "./service";

/**
 * ItemService implements a Service using Item model as
 * type argument
 */
export default class ItemService implements Service<Item> {
    async get(id: number): Promise<Item> {
        const response = await fetch(`http://localhost:8000/items/${id}`, { method: HTTPMethod.GET })

        if (response.status === 200) {
            return await response.json()
        } else {
            throw new Error(`Could not retrieve this todo: ${id}`)
        }
    }

    async getAll(): Promise<Item[]> {
        const response = await fetch('http://localhost:8000/items/', { method: HTTPMethod.GET })

        if (response.status === 200) {
            return await response.json()
        } else {
            throw new Error('Could not retrieve the list of todos')
        }
    }

    async create(item: Item): Promise<Item> {
        const response = await fetch('http://localhost:8000/items/', { method: HTTPMethod.POST, body: JSON.stringify(item) })

        if (response.status === 200) {
            return item
        } else {
            throw new Error('Could not create this item')
        }
    }

    async update(item: Item): Promise<Item> {
        const response = await fetch('http://localhost:8000/items/', { method: HTTPMethod.PUT, body: JSON.stringify(item) })

        if (response.status === 200) {
            return item
        } else {
            throw new Error('Could not update this item')
        }
    }

    async delete(id: number): Promise<number> {
        const response = await fetch(`http://localhost:8000/items/${id}`, { method: HTTPMethod.DELETE })

        if (response.status === 200) {
            return id
        } else {
            throw new Error(`Could not delete this item: ${id}`)
        }
    }
}