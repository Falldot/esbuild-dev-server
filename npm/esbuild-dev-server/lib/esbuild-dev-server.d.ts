declare module "esbuild-dev-server" {
    /**
    * Dev server options.
    * @type {string} Port start local server
    * @type {string} Root html file
    * @type {string} Stacic files
    * @type {string} Working directory
    * @type {() => void} Event reload local server
    */
    export interface DevServerOptions {
        Port: string;
        Index: string;
        StaticDir: string;
        WatchDir: string;
        OnLoad: () => void;
    }
    /**
    * Set dev server options.
    * @param {DevServerOptions} devServerOptions Options.
    */
    export const setOptions: (devServerOptions: DevServerOptions) => void;
    /**
    * Start dev server.
    * @return {Promise<Error | number | null>}
    */
    export const startServer: () => Promise<Error | number | null>;
    /**
    * Send error dev server.
    * @param {string} message Message to dev server.
    * @return {Promise<Error | null>}
    */
    export const sendError: (message: string) => Promise<Error | null>;
    /**
    * Send reload dev server.
    * @return {Promise<Error | null>}
    */
    export const sendReload: () => Promise<Error | null>;
    /**
    * Set dev server options for esbuild.
    * @param {DevServerOptions} devServerOptions Options.
    * @return {Plugin} esbuild plugin.
    */
    export const esBuildDevServer: (options: DevServerOptions) => {
        name: string;
        setup(): void;
    };
}
