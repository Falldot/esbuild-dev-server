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
