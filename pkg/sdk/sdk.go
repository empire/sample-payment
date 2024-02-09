package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

var _ PaymentService = (*Client)(nil)

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
	}
}

func (c *Client) Deposit(accountID AccountID, amount Amount) (string, error) {
	type require struct {
		AccountID AccountID `json:"accountID"`
		Amount    Amount    `json:"amount"`
	}

	reqBody, err := json.Marshal(require{AccountID: accountID, Amount: amount})
	if err != nil {
		return "", err
	}

	httpReq, err := http.NewRequest("POST", c.BaseURL+"/deposit", bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return "", err
	}

	if httpResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("deposit failed: %s", respBody)
	}

	return string(respBody), nil
}

func (c *Client) Balance(accountID AccountID) (Amount, error) {
	type request struct {
		AccountID AccountID `json:"accountID"`
	}

	type response struct {
		Balance Amount `json:"balance"`
	}

	req := request{AccountID: accountID}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return 0, err
	}

	httpReq, err := http.NewRequest("GET", c.BaseURL+"/balance", bytes.NewReader(reqBody))
	if err != nil {
		return 0, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return 0, err
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return 0, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("balance failed: %s", respBody)
	}

	var resp response
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return 0, err
	}
	return resp.Balance, nil
}

func (c *Client) Withdraw(accountID AccountID, amount Amount) (string, error) {
	type request struct {
		AccountID AccountID `json:"accountID"`
		Amount    Amount    `json:"amount"`
	}

	type response struct {
		Message string `json:"message"`
	}

	req := request{AccountID: accountID, Amount: amount}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	httpReq, err := http.NewRequest("POST", c.BaseURL+"/withdraw", bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return "", err
	}

	if httpResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("withdraw failed: %s", respBody)
	}

	var resp response
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}
