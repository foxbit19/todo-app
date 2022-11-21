import { Testing as NSRequest } from "./mockRequest";
import { Testing as NSResponse } from "./mockResponse";

export namespace Testing {
    export class Common {
        public static findQueryString(url: RequestInfo | URL): string {
            const split = url.toString().split('/')
            return split[split.length - 1];
        }

        public static getMockImplementation<T, K>(request: NSRequest.MockRequest<T>, response: NSResponse.MockResponse<K>) {
            return (input: RequestInfo | URL, init?: RequestInit) => {
                // check if the method is correct
                if (init?.method !== request.method) {
                    throw new Error('Wrong HTTP method')
                }

                // check if the queryString is correct, if any
                if (request.queryString && Common.findQueryString(input) !== request.queryString) {
                    throw new Error('Wrong query string')
                }

                // check if the body is correct, if it is present
                if (request.body && init?.body !== JSON.stringify(request.body)) {
                    throw new Error(`Wrong Body provided. got ${init?.body}, want ${JSON.stringify(request.body)}`)
                }

                return Promise.resolve({
                    status: response.status,
                    json: async () => response.body
                } as Response)
            };
        }
    }
}
