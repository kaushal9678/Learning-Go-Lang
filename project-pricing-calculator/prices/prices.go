package prices

import (
	"fmt"

	ioManager "example.com/pricing-calculator/iomanager"
)

type TaxIncludedPriceJob struct{
	IOFileManager ioManager.IOManager `json:"-"` // use - to ignore this key in JSON
	InputPrices     []float64 `json:"input_prices"`
	TaxRate  float64 `json:"tax_rate"`
	TaxIncludedPrices     map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
    strPrices, err := job.IOFileManager.ReadLines()
    if err != nil {
        fmt.Println("Error reading prices:", err)
        return
    }
    var floatPrices []float64
    for _, str := range strPrices {
        var price float64
        if _, err := fmt.Sscanf(str, "%f", &price); err == nil {
            floatPrices = append(floatPrices, price)
        } else {
            fmt.Printf("Invalid price entry: %s\n", str)
        }
    }
    job.InputPrices = floatPrices
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

func NewTaxIncludedPriceJob(iom ioManager.IOManager, taxRate float64) *TaxIncludedPriceJob {
    return &TaxIncludedPriceJob{
        IOFileManager: iom,
        InputPrices:   []float64{},
        TaxRate:       taxRate,
        //TaxIncludedPrices: make(map[string]float64),
    }
}