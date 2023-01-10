package main

func priceCheck(products []string, productPrices []float64, productSold []string, soldPrice []float64) int {
	mapper := make(map[string]float64)
	for i := 0; i < len(products); i++ {
		mapper[products[i]] = productPrices[i]
	}

	count := 0
	for i := 0; i < len(productSold); i++ {
		if val, ok := mapper[productSold[i]]; ok {
			if val != soldPrice[i] {
				count++
			}
		}
	}

	return count
}
