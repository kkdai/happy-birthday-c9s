package main

import (
  "github.com/mbndr/figlet4go"
  "fmt"
)

func main(){

ascii := figlet4go.NewAsciiRender()

// Adding the colors to RenderOptions
options := figlet4go.NewRenderOptions()
options.FontName = "larry3d"
options.FontColor = []figlet4go.Color{
	// Colors can be given by default ansi color codes...
	figlet4go.ColorGreen,
	figlet4go.ColorYellow,
	figlet4go.ColorCyan,
}

// The underscore would be an error
renderStr, _ := ascii.RenderOpts("Happy Birthday c9s", options)
fmt.Print(renderStr)
}
