package devserver

import (
	"github.com/Falldot/esbuild-dev-server/internal/server"
	"github.com/evanw/esbuild/pkg/api"
)

type Options server.DevServerOptions

func Start() error {
	return server.StartDevServer()
}

func Plugin(options Options) api.Plugin {
	server.SetOptions(server.DevServerOptions(options))
	return api.Plugin{
		Name: "dev-server",
		Setup: func(pb api.PluginBuild) {
			pb.OnEnd(func(result *api.BuildResult) {
				if server.IsStart {
					if len(result.Errors) > 0 {
						strs := api.FormatMessages(result.Errors, api.FormatMessagesOptions{
							Kind:  api.ErrorMessage,
							Color: true,
						})
						server.SetError(strs[0])
					} else {
						server.Reload()
					}
				}
			})
		},
	}
}
