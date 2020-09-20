package main

import (
	"fmt"
	"strconv"

	"github.com/tidwall/buntdb"
)

var badgeSVG = "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"%s\" height=\"20\" role=\"img\" aria-label=\"%s: %s\"><title>%s: %s</title><linearGradient id=\"s\" x2=\"0\" y2=\"100%%\"><stop offset=\"0\" stop-color=\"#bbb\" stop-opacity=\".1\"/><stop offset=\"1\" stop-opacity=\".1\"/></linearGradient><clipPath id=\"r\"><rect width=\"%s\" height=\"20\" rx=\"3\" fill=\"#fff\"/></clipPath><g clip-path=\"url(#r)\"><rect width=\"65\" height=\"20\" fill=\"#555\"/><rect x=\"65\" width=\"93\" height=\"20\" fill=\"#%s\"/><rect width=\"%s\" height=\"20\" fill=\"url(#s)\"/></g><g fill=\"#fff\" text-anchor=\"middle\" font-family=\"Verdana,Geneva,DejaVu Sans,sans-serif\" text-rendering=\"geometricPrecision\" font-size=\"110\"><text aria-hidden=\"true\" x=\"335\" y=\"150\" fill=\"#010101\" fill-opacity=\".3\" transform=\"scale(.1)\" textLength=\"550\">%s</text><text x=\"335\" y=\"140\" transform=\"scale(.1)\" fill=\"#fff\" textLength=\"550\">%s</text><text aria-hidden=\"true\" x=\"1105\" y=\"150\" fill=\"#010101\" fill-opacity=\".3\" transform=\"scale(.1)\" textLength=\"830\">%s</text><text x=\"1105\" y=\"140\" transform=\"scale(.1)\" fill=\"#fff\" textLength=\"%s\">%s</text></g></svg>"

func GenerateBadge(title, value, color string) string {
	rectWidth := "100"
	textLength := "200"
	return fmt.Sprintf(badgeSVG, rectWidth, title, value, title, value, rectWidth, color, rectWidth, title, title, value, textLength, value)
}

func set(key string, value int) error {
	err := db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, strconv.Itoa(value), nil)
		return err
	})
	if err != nil {
		fmt.Println("[DEBUG] " + err.Error())
		return err
	}

	return nil
}

func get(key string) (counter int) {
	_ = db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			return err
		}
		if counter, err = strconv.Atoi(val); err != nil {
			fmt.Println("[DEBUG] " + err.Error())
			counter = 0
		}
		return nil
	})

	return counter
}
