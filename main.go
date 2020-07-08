package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("DirPath is needed")
		return
	}

	dupcheck(args[0])
}

func dupcheck(dir string) {
	files := map[string]bool{}
	err := filepath.Walk(dir, visit(files))
	if err != nil {
		log.Fatal(err)
	}
}

func visit(files map[string]bool) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		stat, _ := os.Stat(path)
		if stat.IsDir() {
			return nil
		}

		fn := filepath.Base(path)
		if files[fn] {
			fmt.Println(fn)
		} else {
			files[fn] = true
		}
		return nil
	}
}
