package dto

type SubmitPrintJobRequestBody struct {
	DocumentID    string        `json:"document_id" validate:"required"`
	ClusterID     uint          `json:"cluster_id" validate:"required"`
	PrintSettings PrintSettings `json:"print_settings" validate:"required"`
}

type PrintSettings struct {
	ColorMode   ColorMode   `json:"color_mode" validate:"required"`
	PaperSize   PaperSize   `json:"paper_size" validate:"required"`
	Orientation Orientation `json:"orientation" validate:"required"`
	Copies      int32       `json:"copies" validate:"required"`
	DoubleSided bool        `json:"double_sided" validate:"required"`
}

type ColorMode string

const (
	ColorModeColor     ColorMode = "color"
	ColorModeGreyscale ColorMode = "greyscale"
)

type PaperSize string

const (
	PaperSizeA3 PaperSize = "a3"
	PaperSizeA4 PaperSize = "a4"
	PaperSizeA5 PaperSize = "a5"
	PaperSizeA2 PaperSize = "a2"
)

type Orientation string

const (
	OrientationPortrait  Orientation = "portrait"
	OrientationLandscape Orientation = "landscape"
)

type PrintJobTranfer struct {
	ClusterID     uint
	DocumentID    string
	Name          string
	Content       []byte
	PrintSettings PrintSettings
}
