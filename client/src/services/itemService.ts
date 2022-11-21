import HTTPMethod from "http-method-enum";
import Item from "../models/item";
import Service from "./service";
import ItemAdapter from '../models/itemAdapter';
import { Remote } from './common/common'

/**
 * ItemService implements a Service using Item model as
 * type argument
 */
export default class ItemService implements Service<Item> {
    async get(id: number): Promise<Item> {
        const response = await fetch(`${Remote.Common.getBaseUrl()}/items/${id}`, { method: HTTPMethod.GET })

        if (response.status === 200) {
            return Item.create(await response.json())
        } else {
            throw new Error(`Could not retrieve this todo: ${id}`)
        }
    }

    async getAll(): Promise<Item[]> {
        const response = await fetch(`${Remote.Common.getBaseUrl()}/items/`, { method: HTTPMethod.GET })

        if (response.status === 200) {
            return (await response.json()).map((item: ItemAdapter) => Item.create(item))
        } else {
            throw new Error('Could not retrieve the list of todos')
        }
    }

    async create(item: Item): Promise<Item> {
        const response = await fetch(`${Remote.Common.getBaseUrl()}/items/`, { method: HTTPMethod.POST, body: JSON.stringify(item) })

        if (response.status === 202) {
            return item
        } else {
            throw new Error('Could not create this item')
        }
    }

    async update(item: Item): Promise<Item> {
        const response = await fetch(`${Remote.Common.getBaseUrl()}/items/${item.id}`, { method: HTTPMethod.PUT, body: JSON.stringify(ItemAdapter.create(item)) })

        if (response.status === 202) {
            return item
        } else {
            throw new Error('Could not update this item')
        }
    }

    async delete(id: number): Promise<number> {
        const response = await fetch(`${Remote.Common.getBaseUrl()}/items/${id}`, { method: HTTPMethod.DELETE })

        if (response.status === 200) {
            return id
        } else {
            throw new Error(`Could not delete this item: ${id}`)
        }
    }

    async reorder(sourceId: number, targetId: number): Promise<number> {
        const response = await fetch(`${Remote.Common.getBaseUrl()}/items/reorder/${sourceId}/${targetId}`, { method: HTTPMethod.PATCH })

        if (response.status === 202) {
            return sourceId
        } else {
            throw new Error(`Could not reorder item ${sourceId} with item ${targetId}`)
        }
    }
}