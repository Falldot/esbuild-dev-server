const {build} = require("esbuild")
const esBuildDevServer = require("esbuild-dev-server")

esBuildDevServer.start(
	build({
		entryPoints: ["src/index.js"],
		outdir: "dist/js",
		incremental: true,
		// and more options ...
	}),
	{
		port:      "8080", // optional, default: 8080
		watchDir:  "src", // optional, default: "src"
		index:     "dist/index.html", // optional
		staticDir: "dist", // optional
		onBeforeRebuild: {}, // optional
		onAfterRebuild:  {}, // optional
	}
)