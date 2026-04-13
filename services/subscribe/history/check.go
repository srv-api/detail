package history

func (s *historyService) ExpireTransaction(orderID string) error {
	subscribe, err := s.Repo.FindByOrderID(orderID)
	if err != nil {
		return err
	}

	if subscribe.Status != "pending" {
		return nil // tidak perlu update jika sudah bukan pending
	}

	return s.Repo.UpdateStatus(orderID, "expired")

}
