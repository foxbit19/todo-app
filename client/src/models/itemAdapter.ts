import Item from "./item";
import dayjs from 'dayjs'
export default class ItemAdapter {
    public Id: number;
    public Description: string;
    public Order: number;
    public Completed: boolean;
    public CompletedDate: string;

    constructor(id: number, description: string, order: number, completed: boolean = false, completedDate: string = '') {
        this.Id = id;
        this.Description = description;
        this.Order = order;
        this.Completed = completed
        this.CompletedDate = completedDate
    }

    public static create(item: Item) {
        return new ItemAdapter(item.id, item.description, item.order, item.completed, dayjs(item.completedDate).format('DD MMM YY H:mm ZZ'));
    }
}