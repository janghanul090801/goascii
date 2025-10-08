/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"html/template"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert [image-path]",
	Short: "Convert an image to ASCII art and save it as an HTML file.",
	Long: `Convert an image (JPEG, PNG) to ASCII art and save it as an HTML file.
You can specify the resolution of the ASCII art.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imgPath := args[0]
		resolution, _ := cmd.Flags().GetInt("resolution")
		output, _ := cmd.Flags().GetString("output")

		f, err := os.Open(imgPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			fmt.Println(err)
			return
		}

		bounds := img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()

		ratio := float64(height) / float64(width)

		fontSize := 300.0 / float64(resolution)
		if fontSize < 0.5 {
			fontSize = 0.5
		}

		lineHeight := ratio / 2

		html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>ASCII Art</title>
<style>
	pre {
		font-family: monospace;
		font-size: %.2fpx;
		line-height: %.3f;
	}
</style>
</head>
<body>
<pre>{{.Content}}</pre>
</body>
</html>`, fontSize, lineHeight)

		asciiPixel := "@$#Y!=+~- "
		var index int
		content := ""

		forX := bounds.Max.X / resolution
		forY := bounds.Max.Y / resolution
		if forX == 0 {
			forX = 1
		}
		if forY == 0 {
			forY = 1
		}

		for y := bounds.Min.Y; y < bounds.Max.Y; y += forY {
			for x := bounds.Min.X; x < bounds.Max.X; x += forX {
				r, g, b, a := img.At(x, y).RGBA()
				gray := float64((r*299 + g*587 + b*114) / 1000 >> 8)
				index = int(math.Floor(gray / 25.6))
				if a == 0 {
					index = 9
				}
				content += fmt.Sprintf("%c", asciiPixel[index])
			}
			content += "\n"
		}

		t, err := template.New("page").Parse(html)
		if err != nil {
			fmt.Println(err)
			return
		}

		data := struct {
			Content string
		}{
			content,
		}

		ff, err := os.Create(output)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer ff.Close()

		if err := t.Execute(ff, data); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("ASCII art saved to %s\n", output)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().IntP("resolution", "r", 60, "Resolution of the ASCII art")
	convertCmd.Flags().StringP("output", "o", "templated.html", "Output HTML file name")
}
