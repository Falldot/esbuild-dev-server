package ws

import (
	"html/template"
	"os"
	"strings"
)

const HotReloadScript = `
<script type="text/javascript">
function tryConnectToReload(address) {
	var conn = new WebSocket(address);

	conn.onclose = function() {
		setTimeout(function() {
			tryConnectToReload(address);
		}, 2000);
	};

	conn.onmessage = function(evt) {
		evt.data === "reload" ? location.reload() : console.error(evt.data);
	};
}
try {
    if (window["WebSocket"]) {
        try {
            tryConnectToReload("ws://localhost{port}/connect");
        }
        catch (ex) {
            tryConnectToReload("wss://localhost{port}/connect");
        }
    } else {
        console.log("Your browser does not support WebSockets, cannot connect to the Reload service.");
    }
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
