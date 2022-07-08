package dto

type FilterCarbonProducer struct {
	CategoryCarbonProducerID   string   `query:"category_carbon_producer_id"`
}

type CalculateTransportationCarbonProducer struct {
	JarakTempuh  int `json:"jarak_tempuh" validate:"required"`
	BahanBakar   string `json:"bahan_bakar" validate:"required"`
}

type CalculateElectricityCarbonProducer struct {
	LongUse  int `json:"long_use" validate:"required"`
	NumberOfWatt   int `json:"number_of_watt" validate:"required"`
}