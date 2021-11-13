package ws

import (
	"html/template"
	"os"
	"strings"
)

const HotReloadScript = `
<script type="text/javascript">
	function tryConnectToReload(address) {
		const conn = new WebSocket(address);
		conn.onclose = () => setTimeout(() => tryConnectToReload(address), 2000);
		conn.onmessage = evt => evt.data === "reload" ? location.reload() : console.error(evt.data);
	}
	try {
		window["WebSocket"] ? tryConnectToReload("ws://localhost{port}/connect") : console.log("Your browser does not support WebSockets, cannot connect to the Reload service.");
	} catch (ex) {
		console.error('Exception during connecting to Reload:', ex);
	}
</script>
`

const DefaultHtmlFile = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
	<style>
		body {
			font-family: "Courier New",Courier,mono;
			color: rgb(209, 205, 199);
			background-color: rgb(24, 26, 27);
		}
		pre code {
			display: block;
			overflow: auto;
			white-space: pre;
			font: 14px;
			padding: 7px;
			tab-size: 4;
		}
		.c {
			font-size: 1em;
			font-weight: bold;
		}
	</style>
</head>
<body>
    <h1>Hot reload script</h1>
	<h3><span class="c">
		Insert this script into your index .php or .html file and it will reload every time rebuild your project,<br>
		or in the plugin options specify the path to your .html file, then the script will be added there automatically.
	</h3></span>
	<pre>
		<code>
<span class="c">&lt;script type="text/javascript"&gt;</span>
	<span class="c">function tryConnectToReload(address) {</span>
		<span class="c">const conn = new WebSocket(address);</span>
		<span class="c">conn.onclose = () => setTimeout(() => tryConnectToReload(address), 2000);</span>
		<span class="c">conn.onmessage = evt => evt.data === "reload" ? location.reload() : console.error(evt.data);</span>
	<span class="c">}</span>
	<span class="c">try {</span>
		<span class="c">window["WebSocket"] ? tryConnectToReload("ws://localhost{port}/connect") : console.log("Your browser does not support WebSockets, cannot connect to the Reload service.");</span>
	<span class="c">} catch (ex) {</span>
		<span class="c">console.error('Exception during connecting to Reload:', ex);</span>
	<span class="c">}</span>
<span class="c">&lt;/script&gt;</span>
		</code>
	</pre>
	<h2>Plugin:</h2>
	<h2><a href="https://github.com/Falldot/esbuild-dev-server">esbuild-dev-server</a></h2>
</body>
</html>
`

func GetDefaultHtmlFile(port string) (*template.Template, error) {
	tmpl := template.New("index")
	tmpl, err := tmpl.Parse(strings.Replace(DefaultHtmlFile, "{port}", port, -1))
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func AddHotReloadScript(path string, port string) (*template.Template, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	html := strings.Replace(string(f), "</body>", strings.Replace(HotReloadScript, "{port}", port, -1)+"</body>", -1)
	tmpl := template.New("index")
	tmpl, err = tmpl.Parse(html)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
