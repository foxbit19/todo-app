import Item from "./item";

export default class ItemAdapter {
    public Id: number;
    public Description: string;
    public Order: number;

    constructor(id: number, description: string, order: number) {
        this.Id = id;
        this.Description = description;
        this.Order = order;
    }

    public static create(item: Item) {
        return new ItemAdapter(item.id, item.description, item.order);
    }
}