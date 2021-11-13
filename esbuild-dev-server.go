package devserver

import (
	"log"

	plugin "github.com/Falldot/esbuild-dev-server/internal/api"
	"github.com/evanw/esbuild/pkg/api"
)

type Options struct {
	Port            string
	Index           string
	StaticDir       string
	WatchDir        string
	OnBeforeRebuild func()
	OnAfterRebuild  func()
}

func Start(build api.BuildResult, options Options) {
	if !errorHandler(build.Errors, nil) {
		log.Fatalln("esbuild error!")
	}

	var server plugin.DevServer
	server = plugin.DevServer{
		Port:      options.Port,
		Index:     options.Index,
		StaticDir: options.StaticDir,
		WatchDir:  options.WatchDir,
		OnReload: func() {
			if options.OnBeforeRebuild != nil {
				options.OnBeforeRebuild()
			}
			result := build.Rebuild()
			if errorHandler(result.Errors, server.SendError) {
				server.SendReload()
			}
			if options.OnAfterRebuild != nil {
				options.OnAfterRebuild()
			}
		},
	}

	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}
}

func errorHandler(errors []api.Message, callback func(string)) bool {
	if len(errors) > 0 {
		str := api.FormatMessages(errors, api.FormatMessagesOptions{
			Kind:  api.ErrorMessage,
			Color: true,
		})
		for _, err := range str {
			log.Println(err)
			if callback != nil {
				callback(err)
			}
		}
		return false
	}
	return true
}
