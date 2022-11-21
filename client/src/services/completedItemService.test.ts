import HTTPMethod from "http-method-enum";
import Item from "../models/item";
import ItemAdapter from "../models/itemAdapter";
import { Testing } from "../testing/common";
import CompletedItemService from "./completedItemService";

describe('Completed Item service', () => {
    const date = new Date(2022, 11, 21, 19, 0, 0)
    const itemMock1: Item = new Item(2, 'This is my incredible todo', 3, true, date)

    afterEach(() => {
        jest.restoreAllMocks()
    });

    test('it gets all completed items from the server', async () => {
        jest.spyOn(global, 'fetch').mockImplementation(Testing.Common.getMockImplementation<Item, ItemAdapter[]>(
            {
                method: HTTPMethod.GET
            },
            {
                status: 200,
                body: [ItemAdapter.create(itemMock1)]
            }
        ))
        const service = new CompletedItemService();
        const response = await service.getAll()
        expect(response).toEqual([itemMock1])
    })

    test('it throws an exception on not implemented functions', async () => {
        const service = new CompletedItemService();

        await expect(service.get(1)).rejects.toThrow('Method not implemented.')
        await expect(service.create(itemMock1)).rejects.toThrow('Method not implemented.')
        await expect(service.update(itemMock1)).rejects.toThrow('Method not implemented.')
        await expect(service.delete(1)).rejects.toThrow('Method not implemented.')
    })
})