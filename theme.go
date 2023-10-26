// TODO: replace with go:embed
//go:generate fyne bundle -o bundled.go -package main ./resources/Mplus1-Regular.ttf
//go:generate fyne bundle -append -o bundled.go -package main ./resources/Mplus1-Bold.ttf
//go:generate fyne bundle -append -o bundled.go -package main ./resources/Mplus1Code-Regular.ttf
//go:generate fyne bundle -append -o bundled.go -package main ./resources/Mplus1Code-Bold.ttf

package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type mytheme struct{}

func (m *mytheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	defaultTheme := theme.DefaultTheme()
	switch n {
	default:
		return defaultTheme.Color(n, v)
	case theme.ColorNameDisabled:
		return color.RGBA{R: 0xab, G: 0xab, B: 0xab, A: 0xff}
	}

}

func (m *mytheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return resourceMplus1CodeRegularTtf
	}
	if s.Bold {
		return resourceMplus1BoldTtf
	}
	return resourceMplus1RegularTtf
}

func (m *mytheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m *mytheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	default:
		return theme.DefaultTheme().Size(n)
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNameInnerPadding:
		return 8
	case theme.SizeNameLineSpacing:
		return 4
	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameScrollBar:
		return 16
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameText:
		return 10
	case theme.SizeNameHeadingText:
		return 24
	case theme.SizeNameSubHeadingText:
		return 18
	case theme.SizeNameCaptionText:
		return 11
	case theme.SizeNameInputBorder:
		return 1
	case theme.SizeNameInputRadius:
		return 5
	case theme.SizeNameSelectionRadius:
		return 3
	}
}
