package main

import (
	"github.com/zserge/lorca"
	"fmt"
)

func bind(ui lorca.UI) {
	ui.Bind("profiles", func() ([]string, error) {
		f, err := LoadProfile()
		fmt.Println(f.SectionStrings())
		if err != nil {
			return nil, err
		}
		return f.SectionStrings(), nil
	})
}
