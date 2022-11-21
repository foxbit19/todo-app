import ItemAdapter from "./itemAdapter";
import dayjs from 'dayjs'

/**
 * This class represents a todo item
 */
export default class Item {
    public id: number;
    public description: string;
    public order: number;
    public completed: boolean;
    public completedDate: Date;

    constructor(id: number, description: string, order: number, completed: boolean = false, completedDate: Date = new Date()) {
        this.id = id;
        this.description = description;
        this.order = order;
        this.completed = completed
        this.completedDate = completedDate
    }

    public static create(adapter: ItemAdapter) {
        return new Item(adapter.Id, adapter.Description, adapter.Order, adapter.Completed, dayjs(adapter.CompletedDate).toDate());
    }
}