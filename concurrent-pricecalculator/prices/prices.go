package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate float64 						`json:"tax_rate"`
	InputPrices []float64 					`json:"input_prices"`
	TaxIncludedPrices map[string]string		`json:"tax_included_prices"`
	FileManager iomanager.IOManager			`json:"-"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		InputPrices: []float64{10, 20, 30},
		FileManager: iom,
	}
}

func (t *TaxIncludedPriceJob) Process(doneChan chan bool, errChan chan error) {
	err := t.LoadData()
	if err != nil {
		errChan <- fmt.Errorf("error loading data: %v", err)
		return
	}

	result := make(map[string]string)
	for _, price := range(t.InputPrices) {
		taxIncludedPrice := price * (1 + t.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	
	t.TaxIncludedPrices = result
	err = t.FileManager.WriteResult(t)
	if err != nil {
		errChan <- fmt.Errorf("error writing the result: %v", err)
		return
	}

	doneChan <- true
}

func (t *TaxIncludedPriceJob) LoadData() error{
	lines, err := t.FileManager.ReadLines()
	if err != nil {
		return fmt.Errorf("error reading the file: %v", err)
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return fmt.Errorf("error converting strings to float: %v", err)
	}
	t.InputPrices = prices

	return nil
}