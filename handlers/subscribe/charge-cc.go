package subscribe

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	dto "github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) CardPayment(c echo.Context) error {
	var req dto.CreditCardChargeRequest
	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = userid
	req.CreatedBy = createdBy

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Proses tokenisasi kartu
	transaction, err := h.serviceSubscribe.CardPayment(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, transaction)
}

// handler.go
func (h *domainHandler) TokenizeCardHandler(c echo.Context) error {
	var req dto.TokenizeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid payload"})
	}

	// Bangun URL dengan query parameter
	query := url.Values{}
	query.Set("card_number", req.CardNumber)
	query.Set("card_exp_month", req.ExpiryMonth)
	query.Set("card_exp_year", req.ExpiryYear)
	query.Set("card_cvv", req.CVV)
	query.Set("client_key", os.Getenv("MIDTRANS_CLIENT_KEY"))

	endpoint := dto.GetMidtransTokenize() + query.Encode()

	// Buat GET request dengan query
	httpReq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to build request"})
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to contact Midtrans"})
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if resp.StatusCode != 200 {
		return c.JSON(resp.StatusCode, result)
	}

	// Optional: lanjut ke proses charge di sini kalau mau
	// h.ChargeWithToken(result["token_id"].(string), req.OrderID, req.Amount)

	return c.JSON(http.StatusOK, result)
}
