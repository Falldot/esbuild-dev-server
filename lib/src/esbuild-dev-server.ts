// Copyright (c) 2021 Falldot
// License: MIT License
// MIT License web page: https://opensource.org/licenses/MIT
import { request } from "http";
import { spawn } from "child_process";

/**
* Dev server options.
* @type {string} Port start local server
* @type {string} Root html file
* @type {string} Stacic files
* @type {string} Working directory
* @type {() => void} Event reload local server
*/
export interface DevServerOptions {
    Port:       string,
    Index:      string,
    StaticDir:  string,
    WatchDir:   string,
    OnLoad:     () => void
}

let Options: DevServerOptions;

/**
* Set dev server options.
* @param {DevServerOptions} devServerOptions Options.
*/
export const setOptions = (devServerOptions: DevServerOptions): void => {
    Options = devServerOptions;
};

/**
* Start dev server.
* @return {Promise<Error | number | null>}
*/
export const startServer = (): Promise<Error | number | null> => new Promise((resolve, reject) => {
    const ls = spawn(__dirname + "/../devserver", [":" + Options.Port, Options.Index, Options.StaticDir, Options.WatchDir]);
    ls.stdout.on("data", data => {
        `${data}` === "Reload" ? Options.OnLoad() : console.log(`${data}`);
    });
    ls.stderr.on("data", data => {
        console.log(`${data}`);
    });
    ls.on('error', error => {
        reject(error);
    });
    ls.on("close", code => {
        resolve(code);
    });
});

/**
* Send error dev server.
* @param {string} message Message to dev server.
* @return {Promise<Error | null>}
*/
export const sendError = (message: string): Promise<Error | null> => new Promise((resolve, reject) => {
    const req = request({
        host: 'localhost',
        port: Options.Port,
        path: '/error',
        method: 'POST',
        headers: {
            'Content-Type': 'text/plain'
        }
    }).on("error", err => reject(err));
    const b = Buffer.alloc(message.length);
    b.write(message);
    req.write(b);
    req.end();
    resolve(null);
});

/**
* Send reload dev server.
* @return {Promise<Error | null>}
*/
export const sendReload = (): Promise<Error | null> => new Promise((resolve, reject) => {
    request({
        host: 'localhost',
        port: Options.Port,
        path: '/reload',
        method: 'GET'
    }).on("error", err => reject(err)).end();
    resolve(null);
});

/**
* Set dev server options for esbuild.
* @param {DevServerOptions} devServerOptions Options.
* @return {Plugin} esbuild plugin.
*/
export const esBuildDevServer = (options: DevServerOptions): {
    name: string;
    setup(): void;
} => ({
    name: "dev-server",
    setup() {
        setOptions(options);
    }
});