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

	events := make([]fsnotify.Event, 0)

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				events = append(events, event)

				if len(events) > 1 {
					if events[0] == events[1] && event.Op&fsnotify.Write == fsnotify.Write {
						action()
						events = make([]fsnotify.Event, 0)
						continue
					} else if events[0] != events[1] {
						for _, event := range events {
							if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Remove == fsnotify.Remove || event.Op&fsnotify.Rename == fsnotify.Rename {
								action()
							}
							if event.Op&fsnotify.Create == fsnotify.Create {
								walk(event.Name, watcher)
								events = make([]fsnotify.Event, 0)
							}
						}
					}
				}
				if len(events) > 0 {
					if event.Op&fsnotify.Create == fsnotify.Create {
						walk(event.Name, watcher)
						events = make([]fsnotify.Event, 0)
					}
					if event.Op&fsnotify.Remove == fsnotify.Remove || event.Op&fsnotify.Rename == fsnotify.Rename {
						action()
						events = make([]fsnotify.Event, 0)
					}
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
