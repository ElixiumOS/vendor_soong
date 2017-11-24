package android
type Product_variables struct {
	Has_legacy_camera_hal1 struct {
		Cflags []string
	}

	Uses_media_extensions struct {
		Cflags []string
	}

	Needs_text_relocations struct {
		Cppflags []string
	}

	Cant_reallocate_omx_buffers struct {
		Cflags []string
	}

	Qcom_bsp_legacy struct {
		Cflags []string
	}

	Use_legacy_rescaling struct {
		Cflags []string
	}
}

type ProductVariables struct {
	Has_legacy_camera_hal1  *bool `json:",omitempty"`
	Uses_media_extensions   *bool `json:",omitempty"`
	Needs_text_relocations  *bool `json:",omitempty"`
	QTIAudioPath            *string `json:",omitempty"`
	QTIDisplayPath          *string `json:",omitempty"`
	QTIMediaPath            *string `json:",omitempty"`
	Cant_reallocate_omx_buffers *bool `json:",omitempty"`
	Qcom_bsp_legacy         *bool `json:",omitempty"`
	Use_legacy_rescaling    *bool `json:",omitempty"`
}
