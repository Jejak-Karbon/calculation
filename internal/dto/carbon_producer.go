package dto

type FilterCarbonProducer struct {
	CategoryCarbonProducerID   string   `query:"category_carbon_producer_id"`
}

type CalculateCarbonProducer struct {
	JarakTempuh  int `json:"jarak_tempuh"`
	BahanBakar   string `json:"bahan_bakar"`
	LamaPenggunaan  int `json:"lama_penggunaan"`
	JumlahWatt   int `json:"jumlah_watt"`
}