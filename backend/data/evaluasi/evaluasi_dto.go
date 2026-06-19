package evaluasi

type CreateEvaluasiRequest struct {
	IDMurid           uint   `json:"id_murid" binding:"required"`
	IDPembimbing      uint   `json:"id_pembimbing" binding:"required"`
	EvaluasiKe        uint8  `json:"evaluasi_ke" binding:"required"`
	Nilai             string `json:"nilai" binding:"required"`
	CatatanPembimbing string `json:"catatan_pembimbing"`
	TanggalEvaluasi   string `json:"tanggal_evaluasi" binding:"required"` // YYYY-MM-DD
}

type EvaluasiPerMuridResponse struct {
    IDMurid            uint    `json:"id_murid" gorm:"column:id_murid"`
    KodeMurid          string  `json:"kode_murid" gorm:"column:kode_murid"`
    NamaMurid          string  `json:"nama_murid" gorm:"column:nama_murid"`
    IDPembimbing       *uint   `json:"id_pembimbing,omitempty" gorm:"column:id_pembimbing"`
    NamaPembimbing     *string `json:"nama_pembimbing,omitempty" gorm:"column:nama_pembimbing"`
    Evaluasi1Nilai     *string `json:"evaluasi_1_nilai,omitempty" gorm:"column:evaluasi_1_nilai"`
    Evaluasi1Catatan   *string `json:"evaluasi_1_catatan,omitempty" gorm:"column:evaluasi_1_catatan"`
    Evaluasi1Tanggal   *string `json:"evaluasi_1_tanggal,omitempty" gorm:"column:evaluasi_1_tanggal"`
    Evaluasi2Nilai     *string `json:"evaluasi_2_nilai,omitempty" gorm:"column:evaluasi_2_nilai"`
    Evaluasi2Catatan   *string `json:"evaluasi_2_catatan,omitempty" gorm:"column:evaluasi_2_catatan"`
    Evaluasi2Tanggal   *string `json:"evaluasi_2_tanggal,omitempty" gorm:"column:evaluasi_2_tanggal"`
    Evaluasi3Nilai     *string `json:"evaluasi_3_nilai,omitempty" gorm:"column:evaluasi_3_nilai"`
    Evaluasi3Catatan   *string `json:"evaluasi_3_catatan,omitempty" gorm:"column:evaluasi_3_catatan"`
    Evaluasi3Tanggal   *string `json:"evaluasi_3_tanggal,omitempty" gorm:"column:evaluasi_3_tanggal"`
    TahapBerikutnya    uint8   `json:"tahap_berikutnya"`
    BolehInputTahap2   bool    `json:"boleh_input_tahap_2"`
    BolehInputTahap3   bool    `json:"boleh_input_tahap_3"`
    StatusProgressText string  `json:"status_progress_text"`
}