package main

import (
	"fmt"
	"io/ioutil"
	"context"
	"os"
	"golang.org/x/sync/errgroup"
)

func readFiles(ctx context.Context, files []string) ([]string,error){
	g, ctx := errgroup.WithContext(ctx)

	results := make([]string, len(files))

	for i, file := range files{
		i,file := i,file
		g.Go(func() error {
			data, err := ioutil.ReadFile(file)
			if err == nil {
				results[i] = string(data)
			}
			return err
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

func main(){
	var files = []string{
		"fileRead.txt",
		"questions.txt",
	}

	results, err := readFiles(context.Background(), files)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, result := range results{
		fmt.Println(result)	
	}
}