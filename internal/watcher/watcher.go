package watcher

import (
	"io/fs"
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func Watch(dir string, action func()) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Remove == fsnotify.Remove || event.Op&fsnotify.Rename == fsnotify.Rename {
					action()
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					walk(event.Name, watcher)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	walk(dir, watcher)

	<-done
}

func walk(dir string, watcher *fsnotify.Watcher) {
	if err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return watcher.Add(path)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
