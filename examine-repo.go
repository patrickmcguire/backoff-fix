package main

import (
	"go/parser"
	"go/token"
	"fmt"
)

func ExamineRepository(location string) (err error) {
	fset := token.NewFileSet()
	parsedDir, err := parser.ParseDir(fset, location, nil, 0)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(parsedDir)
	return nil
}
