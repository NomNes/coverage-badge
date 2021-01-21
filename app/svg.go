package app

import "fmt"

const svg = `<svg xmlns="http://www.w3.org/2000/svg" width="112" height="20">
    <linearGradient id="b" x2="0" y2="100%%">
        <stop offset="0" stop-color="#444D56"/>
        <stop offset="1" stop-color="#24292E"/>
    </linearGradient>
    <mask id="a">
        <rect width="112" height="20" rx="3" fill="#fff"/>
    </mask>
    <g mask="url(#a)">
        <path fill="url(#b)" d="M0 0h76v20H0z"/>
        <path fill="#%s" d="M76 0h36v20H76z"/>
    </g>
    <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
        <text x="39" y="15" fill="#010101" fill-opacity=".3">Coverage</text>
        <text x="39" y="14">Coverage</text>
        <text x="93" y="15" fill="#010101" fill-opacity=".3">%s</text>
        <text x="93" y="14">%s</text>
    </g>
</svg>`

func GetSvg(p int) string {
	color := "000000" // black
	value := "n/a"
	if p >= 0 {
		value = fmt.Sprintf("%d%%", p)
		switch {
		case p > 95:
			color = "44cc11" // brightgreen
		case p > 90:
			color = "97ca00" // green
		case p > 75:
			color = "a4a61d" // yellowgreen
		case p > 60:
			color = "dfb317" // yellow
		case p > 40:
			color = "fe7d37" // orange
		default:
			color = "e05d44" // red
		}
	}
	return fmt.Sprintf(svg, color, value, value)
}
