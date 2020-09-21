package main

import (
	"fmt"
	"strconv"

	"github.com/tidwall/buntdb"
	"github.com/valyala/fasttemplate"
)

const badgeSVG = `
<svg
	xmlns="http://www.w3.org/2000/svg"
	xmlns:xlink="http://www.w3.org/1999/xlink" width="{{rectWidth}}" height="20" role="img" aria-label="{{title}}: {{value}}">
	<title>{{title}}: {{value}}</title>
	<linearGradient id="s" x2="0" y2="100%">
		<stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
		<stop offset="1" stop-opacity=".1"/>
	</linearGradient>
	<clipPath id="r">
		<rect width="100" height="20" rx="3" fill="#fff"/>
	</clipPath>
	<g clip-path="url(#r)">
		<rect width="69" height="20" fill="#555"/>
		<rect x="69" width="31" height="20" fill="#97ca00"/>
		<rect width="{{rectWidth}}" height="20" fill="url(#s)"/>
	</g>
	<g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110">
		<text aria-hidden="true" x="355" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{{titleTextLength}}">{{title}}</text>
		<text x="355" y="140" transform="scale(.1)" fill="#fff" textLength="590">{{title}}</text>
		<text aria-hidden="true" x="835" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{{valueTextLength}}">{{value}}</text>
		<text x="835" y="140" transform="scale(.1)" fill="#fff" textLength="{{valueTextLength}}">{{value}}</text>
	</g>
</svg>

`

var t = fasttemplate.New(badgeSVG, "{{", "}}")

func generateBadge(title, value, color string) string {
	rectWidth := "100"
	titleTextLength := "590"

	var valueTextLength string

	if valueLen := len(value); valueLen <= 2 {
		valueTextLength = "130"
	} else {
		valueTextLength = strconv.Itoa(len(value) * 70)
	}

	fmt.Println(valueTextLength)

	return t.ExecuteString(map[string]interface{}{
		"rectWidth":       rectWidth,
		"title":           title,
		"value":           value,
		"color":           color,
		"titleTextLength": titleTextLength,
		"valueTextLength": valueTextLength,
	})
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
