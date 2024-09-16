package main

import (
	"encoding/json"
	"fmt"
	"github.com/p2eengineering/kalp-sdk-public/kalpsdk"
)

// SmartContract provides functions for a calculator contract
type SmartContract struct {
	kalpsdk.Contract
}

// CalculationResult represents the result of an addition
type CalculationResult struct {
	Operand1 float64 `json:"operand1"`
	Operand2 float64 `json:"operand2"`
	Result   float64 `json:"result"`
}

// Add performs addition of two numbers and stores the result in the world state
func (s *SmartContract) Add(ctx kalpsdk.TransactionContextInterface, calculationID string, operand1 float64, operand2 float64) error {
	result := operand1 + operand2

	calculationResult := CalculationResult{
		Operand1: operand1,
		Operand2: operand2,
		Result:   result,
	}

	resultAsBytes, err := json.Marshal(calculationResult)
	if err != nil {
		return fmt.Errorf("failed to marshal calculation result: %s", err.Error())
	}

	return ctx.PutStateWithoutKYC(calculationID, resultAsBytes)
}

// GetCalculationResult retrieves the calculation result from the world state using the given ID
func (s *SmartContract) GetCalculationResult(ctx kalpsdk.TransactionContextInterface, calculationID string) (*CalculationResult, error) {
	resultAsBytes, err := ctx.GetState(calculationID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %s", err.Error())
	}

	if resultAsBytes == nil {
		return nil, fmt.Errorf("calculation ID %s does not exist", calculationID)
	}

	var calculationResult CalculationResult
	err = json.Unmarshal(resultAsBytes, &calculationResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal calculation result: %s", err.Error())
	}

	return &calculationResult, nil
}