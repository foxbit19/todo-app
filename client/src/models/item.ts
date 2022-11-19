import ItemAdapter from "./itemAdapter";

/**
 * This class represents a todo item
 */
export default class Item {
    public id: number;
    public description: string;
    public order: number;

    constructor(id: number, description: string, order: number) {
        this.id = id;
        this.description = description;
        this.order = order;
    }

    public static create(adapter: ItemAdapter) {
        return new Item(adapter.Id, adapter.Description, adapter.Order);
    }
}