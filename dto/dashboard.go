package dto

import (
	"encoding/json"
	"strconv"
	"strings"
)

type GetDashboardRequest struct {
	ID                    string `param:"id" validate:"required"`
	MerchantID            string `json:"merchant_id"`
	CreatedBy             string `json:"created_by"`
	TotalWebOrder         int64  `json:"total_order_web"`
	TotalProductsActive   int64  `json:"total_products_active"`
	TotalProductsInactive int64  `json:"total_products_inactive"`
	TotalPrice            int    `json:"total_price"`
	TotalWaiting          int    `json:"total_waiting"`
	PaymentCancel         int64  `json:"payment_cancel"`
}

type GetDashboardResponse struct {
	MerchantName          string              `json:"merchant_name"`
	TotalWebOrder         int64               `json:"total_web_order"`
	TotalProductsActive   int64               `json:"total_products_active"`
	TotalProductsInactive int64               `json:"total_products_inactive"`
	TotalSales            int64               `json:"total_sales"`
	TotalPrice            int                 `json:"total_price"`
	WaitingPayment        int64               `json:"total_waiting_payment"`
	TotalWaiting          int                 `json:"total_waiting"`
	PaymentCancel         int64               `json:"payment_cancel"`
	ProductPercentages    []ProductPercentage `json:"product_percentages"`
	IsSubscribed          bool                `json:"is_subscribed"`
}

type ProductPercentage struct {
	ProductID         string  `json:"product_id"`
	ProductName       string  `json:"product_name"`
	TotalLunas        int64   `json:"total_lunas"`
	TotalTransactions int64   `json:"total_transactions"`
	Percentage        float64 `json:"percentage"`
}

func (r GetDashboardResponse) MarshalJSON() ([]byte, error) {
	type Alias GetDashboardResponse
	return json.Marshal(&struct {
		*Alias
		TotalPrice   string `json:"total_price"`
		TotalWaiting string `json:"total_waiting"`
	}{
		Alias:        (*Alias)(&r),
		TotalPrice:   formatRupiah(r.TotalPrice),
		TotalWaiting: formatRupiah(r.TotalWaiting),
	})
}

func formatRupiah(amount int) string {
	s := strconv.Itoa(amount)
	var result strings.Builder
	length := len(s)
	count := 0

	for i := length - 1; i >= 0; i-- {
		if count > 0 && count%3 == 0 {
			result.WriteString(".")
		}
		result.WriteByte(s[i])
		count++
	}

	runes := []rune(result.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return "Rp " + string(runes)
}
