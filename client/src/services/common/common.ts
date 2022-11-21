export namespace Remote {
    export class Common {
        /**
         * Gets the base url
         * @returns a string of the base url to use in this service
         */
        public static getBaseUrl() {
            return process.env.REACT_APP_BASE_URL;
        }
    }
}