package dto

type FilterCarbonProducer struct {
	CategoryCarbonProducerID   string   `query:"category_carbon_producer_id"`
}

type CalculateCarbonProducer struct {
	JarakTempuh  int `json:"jarak_tempuh" validate:"required"`
	BahanBakar   string `json:"bahan_bakar" validate:"required"`
}