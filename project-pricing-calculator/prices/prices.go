package prices

import (
	"fmt"

	filemanager "example.com/pricing-calculator/fileManager"
)

type TaxIncludedPriceJob struct{
	IOFileManager *filemanager.FileManager `json:"-"` // use - to ignore this key in JSON
	InputPrices     []float64 `json:"input_prices"`
	TaxRate  float64 `json:"tax_rate"`
	TaxIncludedPrices     map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData(){
	prices, err := job.IOFileManager.ReadLines()
	if err != nil {
		fmt.Println("Error reading prices:", err)
		return
	}
	job.InputPrices = prices;

	
}
func (job *TaxIncludedPriceJob) Process(){
	job.LoadData();
	result := make(map[string]string)
		for _, price := range job.InputPrices {
			taxIncludedPrice := price *(1 + job.TaxRate) 
			result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
		}
		//fmt.Println(result)
		job.TaxIncludedPrices = result;
		job.IOFileManager.WriteJSON( job)
	}

func NewTaxIncludedPriceJob(fm *filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
    return &TaxIncludedPriceJob{
        IOFileManager: fm,
        InputPrices:   []float64{},
        TaxRate:       taxRate,
        //TaxIncludedPrices: make(map[string]float64),
    }
}