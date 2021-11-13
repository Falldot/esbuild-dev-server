// Copyright (c) 2021 Falldot
// License: MIT License
// MIT License web page: https://opensource.org/licenses/MIT
import { request } from "http";
import { spawn } from "child_process";
import os = require('os');

/**
* Dev server options.
* @type {string} port start local server
* @type {string} Root html file
* @type {string} Stacic files
* @type {string} Working directory
* @type {() => void} Event before rebuild
* @type {() => void} Event after rebuild
*/
export interface DevServerOptions {
    port:       string,
    index:      string,
    staticDir:  string,
    watchDir:   string,
    onBeforeRebuild: () => void
    onAfterRebuild:  () => void
}

let Options: DevServerOptions;

/**
* Start dev server.
* @param {BuildIncremental} result esbuild context.
* @return {Promise<Error | number | null>}
*/
const startServer = (result: any): Promise<Error | number | null> => new Promise((resolve, reject) => {
    const platform = `esbuild-dev-server-${process.platform}-${os.arch()}`;
    const ls = spawn(__dirname + `/../../${platform}/devserver`, ['-p', Options.port, '-i', Options.index, '-s', Options.staticDir, '-w', Options.watchDir]);
    ls.stdout.on("data", data => {
        if (`${data}` === "Reload") {
                result.rebuild().then(() => {
                    sendReload();
                }).catch((err: any) => {
                    sendError(err.message);
                })
        } else {
            console.log(`${data}`);
        }
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
        port: Options.port,
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
        port: Options.port,
        path: '/reload',
        method: 'GET'
    }).on("error", err => reject(err)).end();
    resolve(null);
});

/**
* Set dev server options for esbuild.
* @param {Promise<BuildIncremental>} build esbuild options.
* @param {Promise<DevServerOptions>} options dev-server options.
*/
export const start = (build: Promise<any>, options: DevServerOptions) => {
    build.then((result: any) => {
        Options = options;
        startServer(result)
    }).catch((err: any) => {
        console.log(err)
    })
}