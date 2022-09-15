package main

import (
	"fmt"
	"os"
	"strings"
)

var heartBeatTpl = "<?xml version=\"1.0\" encoding=\"utf-8\" ?>\n<svg baseProfile=\"full\" height=\"150\" version=\"1.1\" viewBox=\"0,0,200,200\" width=\"150\"\n     xmlns=\"http://www.w3.org/2000/svg\" xmlns:ev=\"http://www.w3.org/2001/xml-events\"\n     xmlns:xlink=\"http://www.w3.org/1999/xlink\"><defs />\n    <g transform=\"translate(100 100)\">\n        <path d=\"M92.71,7.27L92.71,7.27c-9.71-9.69-25.46-9.69-35.18,0L50,14.79l-7.54-7.52C32.75-2.42,17-2.42,7.29,7.27v0 c-9.71,9.69-9.71,25.41,0,35.1L50,85l42.71-42.63C102.43,32.68,102.43,16.96,92.71,7.27z\" fill=\"tomato\" stroke=\"#B6BBC1\" stroke-width=\"2\" transform=\"translate(-50 -50)\" />\n        <animateTransform additive=\"sum\" attributeName=\"transform\" dur=\"1.2s\" repeatCount=\"indefinite\" type=\"scale\" values=\"1; 1.5; 1.25; 1.5; 1.5; 1;\" />\n%s\n    </g>\n</svg>"

func main() {
	if len(os.Args) != 3 {
		os.Stdout.Write([]byte("args error!"))
		return
	}
	values := strings.Split(os.Args[2], "\n")

	tpl := "        <text fill=\"bisque\" style=\"font-size:25px; font-family:Arial\" transform=\"translate(-12.5 0)\">%s<animate attributeName=\"visibility\" dur=\"12.0s\" keyTimes=\"%s\" repeatCount=\"indefinite\" values=\"%s\" /></text>\n"
	f, _ := os.Create("assets/heart.svg")

	value := ""
	valuesLength := len(values)
	hiddenSlice := make([]string, valuesLength)
	keyTimes := make([]string, valuesLength+1)
	step := 1.0 / float64(valuesLength)
	for i := 0; i < valuesLength; i++ {
		hiddenSlice[i] = "hidden"
		keyTimes[i] = fmt.Sprintf("%.3f", float64(i)*step)
	}
	keyTimes[valuesLength] = "1"
	for i := 0; i < valuesLength; i++ {
		tmpHidden := make([]string, valuesLength)
		copy(tmpHidden, hiddenSlice)
		tmpHidden[i] = "visible"
		tmpHidden = append(tmpHidden, "hidden")
		value += fmt.Sprintf(tpl, values[i], strings.Join(keyTimes, ";"), strings.Join(tmpHidden, ";"))
	}
	_, err := f.Write([]byte(fmt.Sprintf(heartBeatTpl, value)))
	if err != nil {
		os.Stdout.Write([]byte(err.Error()))
		return
	}
}
