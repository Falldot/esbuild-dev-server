const {build, formatMessages} = require("esbuild")
const {esBuildDevServer, startServer, sendError, sendReload} = require("esbuild-dev-server")

;(async () => {
	const builder = await build({
		bundle: true,
		format: "iife",
		define: { "process.env.NODE_ENV": JSON.stringify(process.env.NODE_ENV || "development") },
		entryPoints: ["src/index.js"],
		incremental: true,
		minify: false,
		sourcemap: true,
		plugins: [
			esBuildDevServer({
				Port: "8080",
				Index: "dist/index.html",
				StaticDir: "dist",
				WatchDir: "src",
				OnLoad: async () => {
					try {
						await builder.rebuild();
						await sendReload();
					} catch(result) {
						let str = await formatMessages(result.errors, {kind: 'error', color: true})
						await sendError(str.join(""));
					}
				}
			})
		],
		target: 'chrome90',
		outdir: "dist/js",
	})
	await startServer()
})()