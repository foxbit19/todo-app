import HTTPMethod from 'http-method-enum';
import Item from '../models/item';
import ItemAdapter from '../models/itemAdapter';
import { Testing } from '../testing/common';
import ItemService from './itemService';

describe('Item service', () => {
    const itemMock1: Item = new Item(1, 'This is my beautiful todo', 1)
    const itemMock2: Item = new Item(2, 'This is my incredible todo', 3)

    afterEach(() => {
        jest.restoreAllMocks()
    });

    describe('gets an item', () => {
        test('it gets an item from the server', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(Testing.Common.getMockImplementation<Item, ItemAdapter>(
                {
                    method: HTTPMethod.GET,
                    queryString: '1'
                },
                {
                    status: 200,
                    body: ItemAdapter.create(itemMock1)
                }
            ))

            const service = new ItemService();
            const response = await service.get(1)
            expect(response.id).toEqual(itemMock1.id)
        })

        test('it throws an error if response status is different from 200', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(Testing.Common.getMockImplementation<Item, ItemAdapter>(
                {
                    method: HTTPMethod.GET,
                    queryString: '1'
                },
                {
                    status: 400,
                }
            ))

            const service = new ItemService();

            await expect(service.get(1)).rejects.toThrowError('Could not retrieve this todo: 1');
        })
    })

    describe('get all items', () => {
        test('it gets all items from the server', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(Testing.Common.getMockImplementation<Item, ItemAdapter[]>(
                {
                    method: HTTPMethod.GET
                },
                {
                    status: 200,
                    body: [ItemAdapter.create(itemMock1), ItemAdapter.create(itemMock2)]
                }
            ))
            const service = new ItemService();
            const response = await service.getAll()

            const want = [itemMock1, itemMock2]

            for (let i = 0; i < response.length; i++) {
                expect(response[i].id).toEqual(want[i].id)
            }
        })

        test('it throws an error if response status is different from 200', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(Testing.Common.getMockImplementation<Item, Item[]>(
                {
                    method: HTTPMethod.GET
                },
                {
                    status: 400,
                }
            ))

            const service = new ItemService();

            await expect(service.getAll()).rejects.toThrowError('Could not retrieve the list of todos');
        })
    })

    describe('create an item', () => {
        test('it creates a new item', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(Testing.Common.getMockImplementation<Item, Item>(
                {
                    method: HTTPMethod.POST,
                    body: itemMock1
                },
                {
                    status: 202,
                }
            ))

            const service = new ItemService();
            const response = await service.create(itemMock1)
            expect(response).toEqual(itemMock1)
        })


        test('it throws an error if response status is different from 200', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(Testing.Common.getMockImplementation<Item, Item[]>(
                {
                    method: HTTPMethod.POST
                },
                {
                    status: 400,
                }
            ))

            const service = new ItemService();
            await expect(service.create(itemMock1)).rejects.toThrowError('Could not create this item');
        })
    })

    describe('update an item', () => {
        test('it updates a new item', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(Testing.Common.getMockImplementation<ItemAdapter, Item>(
                {
                    method: HTTPMethod.PUT,
                    queryString: itemMock1.id.toString(),
                    body: ItemAdapter.create(itemMock1),
                },
                {
                    status: 202,
                }
            ))

            const service = new ItemService();
            const response = await service.update(itemMock1)
            expect(response).toEqual(itemMock1)
        })


        test('it throws an error if response status is different from 200', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(Testing.Common.getMockImplementation<Item, Item[]>(
                {
                    method: HTTPMethod.PUT,
                    queryString: itemMock1.id.toString()
                },
                {
                    status: 400,
                }
            ))

            const service = new ItemService();
            await expect(service.update(itemMock1)).rejects.toThrowError('Could not update this item');
        })
    })

    describe('delete an item', () => {
        test('it delets a new item', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(Testing.Common.getMockImplementation<Item, Item>(
                {
                    method: HTTPMethod.DELETE,
                    queryString: '1'
                },
                {
                    status: 200,
                }
            ))

            const service = new ItemService();
            const response = await service.delete(1)
            expect(response).toEqual(1)
        })


        test('it throws an error if response status is different from 200', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(Testing.Common.getMockImplementation<Item, Item[]>(
                {
                    method: HTTPMethod.DELETE,
                    queryString: '1'
                },
                {
                    status: 400,
                }
            ))

            const service = new ItemService();
            await expect(service.delete(1)).rejects.toThrowError('Could not delete this item: 1');
        })
    })

    describe('reorder an item', () => {
        test('it rises the priority of an item', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(Testing.Common.getMockImplementation<Item, Item>(
                {
                    method: HTTPMethod.PATCH,
                },
                {
                    status: 202,
                }
            ))

            const service = new ItemService();
            const response = await service.reorder(2, 1)
            expect(response).toEqual(2)
        })
    })
})