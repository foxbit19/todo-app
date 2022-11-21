/**
 * This interface defines how a service must
 * be constructed to implements correctly a CRUD logic.
 * It uses generics so it requires a concrete type to construct the service.
 */
export default interface Service<T> {
    get(id: number): Promise<T>
    getAll(): Promise<T[]>
    create(item: T): Promise<T>
    update(item: T): Promise<T>
    delete(id: number): Promise<number>
}