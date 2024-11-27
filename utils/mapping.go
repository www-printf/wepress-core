package utils

import (
	"github.com/www-printf/wepress-core/modules/printer/dto"
	"github.com/www-printf/wepress-core/modules/printer/proto"
)

func MapColorMode(mode dto.ColorMode) proto.PrintSettings_ColorMode {
	switch mode {
	case dto.ColorModeColor:
		return proto.PrintSettings_COLOR_MODE_COLOR
	case dto.ColorModeGreyscale:
		return proto.PrintSettings_COLOR_MODE_GRAYSCALE
	default:
		return proto.PrintSettings_COLOR_MODE_COLOR_UNSPECIFIED
	}
}

func MapPaperSize(size dto.PaperSize) proto.PrintSettings_PaperSize {
	switch size {
	case dto.PaperSizeA3:
		return proto.PrintSettings_PAPER_SIZE_A3
	case dto.PaperSizeA4:
		return proto.PrintSettings_PAPER_SIZE_A4
	case dto.PaperSizeA5:
		return proto.PrintSettings_PAPER_SIZE_A5
	case dto.PaperSizeA2:
		return proto.PrintSettings_PAPER_SIZE_A2
	default:
		return proto.PrintSettings_PAPER_SIZE_SIZE_UNSPECIFIED
	}
}

func MapOrientation(orientation dto.Orientation) proto.PrintSettings_Orientation {
	switch orientation {
	case dto.OrientationPortrait:
		return proto.PrintSettings_ORIENTATION_PORTRAIT
	case dto.OrientationLandscape:
		return proto.PrintSettings_ORIENTATION_LANDSCAPE
	default:
		return proto.PrintSettings_ORIENTATION_UNSPECIFIED
	}
}
