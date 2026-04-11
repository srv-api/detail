package subscribe

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"time"

	dto "github.com/srv-api/merchant/dto"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) ChargeShopeePay(c echo.Context) error {
	var req dto.ChargeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
	}

	if req.OrderID == "" || req.GrossAmount <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Missing required fields: order_id or amount",
		})
	}

	payload := map[string]interface{}{
		"payment_type": "shopeepay",
		"transaction_details": map[string]interface{}{
			"order_id":     req.OrderID,
			"gross_amount": req.GrossAmount,
		},
		"shopeepay": map[string]interface{}{
			"callback_url": "https://your-callback-url.com/notification", // Sesuaikan
			"redirect_url": "https://your-web-app.com/payment/success",   // Jika ingin redirect ke frontend
		},
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to encode payload",
		})
	}

	httpReq, err := http.NewRequest("POST", dto.GetMidtransEndpoint(), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create HTTP request",
		})
	}

	auth := base64.StdEncoding.EncodeToString([]byte(dto.GetMidtransServerKey() + ":"))
	httpReq.Header.Set("Authorization", "Basic "+auth)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 20 * time.Second}
	res, err := client.Do(httpReq)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"error":   "Failed to contact Midtrans",
			"details": err.Error(),
		})
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to read Midtrans response",
		})
	}

	var parsed dto.ShopeePayResponse
	if err := json.Unmarshal(resBody, &parsed); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   "Invalid response from Midtrans",
			"details": string(resBody),
		})
	}

	if parsed.StatusCode != "201" {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"error":   "Midtrans returned an error",
			"code":    parsed.StatusCode,
			"message": parsed.StatusMessage,
		})
	}

	return c.JSON(http.StatusOK, parsed)
}
