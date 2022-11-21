export namespace Testing {
    /**
     * A support interface to collect response argument
     * for testing purposes
     */
    export interface MockResponse<T> {
        status: number,
        body?: T
    }
}