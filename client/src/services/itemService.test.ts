import HTTPMethod from 'http-method-enum';
import Item from '../models/item';
import ItemAdapter from '../models/itemAdapter';
import ItemService from './itemService';

/**
 * A support interface to collect request argument
 * for testing purposes
 */
interface MockRequest<T> {
    method: string,
    queryString?: string,
    body?: T
}

/**
 * A support interface to collect response argument
 * for testing purposes
 */
interface MockResponse<T> {
    status: number,
    body?: T
}

describe('Item service', () => {
    const itemMock1: Item = new Item(1, 'This is my beautiful todo', 1)
    const itemMock2: Item = new Item(2, 'This is my incredible todo', 3)

    const findQueryString = (url: RequestInfo | URL): string => {
        const split = url.toString().split('/')
        return split[split.length - 1];
    }

    const getMockImplementation = <T, K>(request: MockRequest<T>, response: MockResponse<K>) => {
        return (input: RequestInfo | URL, init?: RequestInit) => {
            // check if the method is correct
            if (init?.method !== request.method) {
                throw new Error('Wrong HTTP method')
            }

            // check if the queryString is correct, if any
            if (request.queryString && findQueryString(input) !== request.queryString) {
                throw new Error('Wrong query string')
            }

            // check if the body is correct, if it is present
            if (request.body && init?.body !== JSON.stringify(request.body)) {
                throw new Error('Wrong Body provided')
            }

            return Promise.resolve({
                status: response.status,
                json: async () => response.body
            } as Response)
        };
    }

    afterEach(() => {
        jest.restoreAllMocks()
    });

    describe('gets an item', () => {
        test('it gets an item from the server', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(getMockImplementation<Item, ItemAdapter>(
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
            expect(response).toEqual(itemMock1)
        })

        test('it throws an error if response status is different from 200', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(getMockImplementation<Item, ItemAdapter>(
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
            jest.spyOn(global, 'fetch').mockImplementation(getMockImplementation<Item, ItemAdapter[]>(
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
            expect(response).toEqual([itemMock1, itemMock2])
        })


        test('it throws an error if response status is different from 200', async () => {
            jest.spyOn(global, 'fetch').mockImplementation(getMockImplementation<Item, Item[]>(
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
            jest.spyOn(global, 'fetch').mockImplementation(getMockImplementation<Item, Item>(
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
            jest.spyOn(global, 'fetch').mockImplementation(getMockImplementation<Item, Item[]>(
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
            jest.spyOn(global, 'fetch').mockImplementation(getMockImplementation<Item, Item>(
                {
                    method: HTTPMethod.PUT,
                    queryString: itemMock1.id.toString(),
                    body: itemMock1,
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
            jest.spyOn(global, 'fetch').mockImplementation(getMockImplementation<Item, Item[]>(
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
            jest.spyOn(global, 'fetch').mockImplementation(getMockImplementation<Item, Item>(
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
            jest.spyOn(global, 'fetch').mockImplementation(getMockImplementation<Item, Item[]>(
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
})