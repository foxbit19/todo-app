export namespace Testing {
    /**
     * A support interface to collect request argument
     * for testing purposes
     */
    export interface MockRequest<T> {
        method: string,
        queryString?: string,
        body?: T
    }
}