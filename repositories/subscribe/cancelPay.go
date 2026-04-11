package subscribe

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *subscribeRepository) CancelPay(request dto.GetorderID) ([]byte, int, error) {
	baseURL := dto.GetMidtransUrl()
	serverKey := dto.GetMidtransServerKey()

	url := fmt.Sprintf("%s/%s/cancel", baseURL, request.OrderID)
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(serverKey+":"))

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, nil
}

func (r *subscribeRepository) UpdateSubscribeByOrderID(data dto.MidtransCancelResponse) error {
	var subscribe entity.Subscribe

	// cari berdasarkan order_id
	if err := r.DB.Where("order_id = ?", data.OrderID).First(&subscribe).Error; err != nil {
		return err
	}

	// parse TransactionTime
	transactionTime, err := time.Parse("2006-01-02 15:04:05", data.TransactionTime)
	if err != nil {
		return err
	}

	// convert GrossAmount string ke int64
	amount, _ := strconv.ParseInt(strings.Split(data.GrossAmount, ".")[0], 10, 64)

	// update field
	subscribe.Status = data.TransactionStatus
	subscribe.TransactionID = data.TransactionID
	subscribe.TransactionTime = transactionTime
	subscribe.GrossAmount = amount
	subscribe.PaymentType = data.PaymentType
	subscribe.UpdatedAt = time.Now()

	return r.DB.Save(&subscribe).Error
}
