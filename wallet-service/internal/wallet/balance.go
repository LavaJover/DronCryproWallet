package wallet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

type TronAccountResponse struct {
	Data []struct {
		Balance  int64 `json:"balance"`  // Основной баланс TRX
		Tokens   []struct {
			TokenID string `json:"tokenId"`
			Amount  int64  `json:"amount"`
		} `json:"trc20token_balances"` // Баланс токенов TRC-20
	} `json:"data"`
}

func GetUSDTBalance(walletAddress string, apiKey string) (float64, error) {
    url := fmt.Sprintf("https://api.trongrid.io/v1/accounts/%s", walletAddress)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return 0, fmt.Errorf("failed to create request: %v", err)
    }
    req.Header.Set("TRON-PRO-API-KEY", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return 0, fmt.Errorf("failed to fetch balance: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return 0, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
    }

    var response TronAccountResponse
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return 0, fmt.Errorf("failed to decode response: %v", err)
    }

    // Проверяем, есть ли данные
    if len(response.Data) == 0 {
        return 0, fmt.Errorf("no data found for address: %s", walletAddress)
    }

	return float64(response.Data[0].Balance) / 1e6, nil
}
