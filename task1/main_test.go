package main

import "testing"

func Test_priceCheck(t *testing.T) {
	type args struct {
		products      []string
		productPrices []float64
		productSold   []string
		soldPrice     []float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Default case: ",
			args: args{
				products:      []string{"eggs", "milk", "cheese"},
				productPrices: []float64{2.89, 3.29, 5.79},
				productSold:   []string{"eggs", "eggs", "cheese", "milk"},
				soldPrice:     []float64{2.89, 2.99, 5.97, 3.29},
			},
			want: 2,
		},
		{
			name: "Custom case: ",
			args: args{
				products:      []string{"eggs", "milk", "cheese", "strawberry", "water"},
				productPrices: []float64{2.89, 3.29, 5.79, 6.14, 0.45},
				productSold:   []string{"eggs", "eggs", "cheese", "milk", "water", "water", "strawberry", "cheese", "milk"},
				soldPrice:     []float64{2.89, 2.99, 5.97, 3.29, 0.45, 0.55, 6.13, 5.79, 3.3},
			},
			want: 5,
		},
		{
			name: "Nothing sold case: ",
			args: args{
				products:      []string{"eggs", "milk", "cheese"},
				productPrices: []float64{2.89, 3.29, 5.79},
				productSold:   []string{},
				soldPrice:     []float64{},
			},
			want: 0,
		},
		{
			name: "Empty products case: ",
			args: args{
				products:      []string{},
				productPrices: []float64{},
				productSold:   []string{"eggs", "eggs", "cheese", "milk"},
				soldPrice:     []float64{2.89, 2.99, 5.97, 3.29},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := priceCheck(tt.args.products, tt.args.productPrices, tt.args.productSold, tt.args.soldPrice); got != tt.want {
				t.Errorf("priceCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
