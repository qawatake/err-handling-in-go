package main

import "fmt"

func main() {
	result, err := Purchase("P-123", 5)
	if err != nil {
		panic(err)
	}
	remaining, e := result.Or()
	if e != nil {
		fmt.Printf("購入できません。在庫は残り%d個ですが、%d個必要です。\n", e.Available, e.Requested)
		return
	}
	fmt.Printf("購入できました。残り%d個\n", remaining)
}

func Purchase(productID string, amount int) (*PurchaseResult, error) {
	stock, err := GetStock(productID)
	if err != nil {
		return nil, fmt.Errorf("在庫取得失敗: %w", err) // 技術的例外
	}
	if amount > stock {
		return NewErr[PurchaseResult](OutOfStockErr{
			Requested: amount,
			Available: stock,
		}), nil // ビジネス例外
	}
	remaining := stock - amount
	err = SaveStock(productID, remaining)
	if err != nil {
		return nil, fmt.Errorf("在庫保存失敗: %w", err) // 技術的例外
	}
	return NewValue[PurchaseResult](remaining), nil // 成功
}

type PurchaseResult = Result[int, OutOfStockErr]

type OutOfStockErr struct {
	Requested int
	Available int
}
