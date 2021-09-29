var __create = Object.create;
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __getProtoOf = Object.getPrototypeOf;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __markAsModule = (target) => __defProp(target, "__esModule", { value: true });
var __export = (target, all) => {
  __markAsModule(target);
  for (var name in all)
    __defProp(target, name, { get: all[name], enumerable: true });
};
var __reExport = (target, module2, desc) => {
  if (module2 && typeof module2 === "object" || typeof module2 === "function") {
    for (let key of __getOwnPropNames(module2))
      if (!__hasOwnProp.call(target, key) && key !== "default")
        __defProp(target, key, { get: () => module2[key], enumerable: !(desc = __getOwnPropDesc(module2, key)) || desc.enumerable });
  }
  return target;
};
var __toModule = (module2) => {
  return __reExport(__markAsModule(__defProp(module2 != null ? __create(__getProtoOf(module2)) : {}, "default", module2 && module2.__esModule && "default" in module2 ? { get: () => module2.default, enumerable: true } : { value: module2, enumerable: true })), module2);
};

// lib/src/esbuild-dev-server.ts
__export(exports, {
  esBuildDevServer: () => esBuildDevServer,
  sendError: () => sendError,
  sendReload: () => sendReload,
  setOptions: () => setOptions,
  startServer: () => startServer
});
var import_http = __toModule(require("http"));
var import_child_process = __toModule(require("child_process"));
var os = require("os");
var Options;
var setOptions = (devServerOptions) => {
  Options = devServerOptions;
};
var startServer = () => new Promise((resolve, reject) => {
  const platform = `esbuild-dev-server-${process.platform}-${os.arch()}`;
  const ls = (0, import_child_process.spawn)(__dirname + `/../../${platform}/devserver`, [":" + Options.Port, Options.Index, Options.StaticDir, Options.WatchDir]);
  ls.stdout.on("data", (data) => {
    `${data}` === "Reload" ? Options.OnLoad() : console.log(`${data}`);
  });
  ls.stderr.on("data", (data) => {
    console.log(`${data}`);
  });
  ls.on("error", (error) => {
    reject(error);
  });
  ls.on("close", (code) => {
    resolve(code);
  });
});
var sendError = (message) => new Promise((resolve, reject) => {
  const req = (0, import_http.request)({
    host: "localhost",
    port: Options.Port,
    path: "/error",
    method: "POST",
    headers: {
      "Content-Type": "text/plain"
    }
  }).on("error", (err) => reject(err));
  const b = Buffer.alloc(message.length);
  b.write(message);
  req.write(b);
  req.end();
  resolve(null);
});
var sendReload = () => new Promise((resolve, reject) => {
  (0, import_http.request)({
    host: "localhost",
    port: Options.Port,
    path: "/reload",
    method: "GET"
  }).on("error", (err) => reject(err)).end();
  resolve(null);
});
var esBuildDevServer = (options) => ({
  name: "dev-server",
  setup() {
    setOptions(options);
  }
});
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  esBuildDevServer,
  sendError,
  sendReload,
  setOptions,
  startServer
});
