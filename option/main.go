package main

import "fmt"

func main() {
	result, err := Purchase("P-123", 5)
	if err != nil {
		panic(err)
	}
	remaining, ok := result.Remaining.V()
	if !ok {
		if outOfStock, ok := result.OutOfStock.V(); ok {
			fmt.Printf("購入できません。在庫は残り%d個ですが、%d個必要です。\n", outOfStock.Available, outOfStock.Requested)
			return
		} else {
			panic("unreachable")
		}
	}
	fmt.Printf("残り%d個\n", remaining)
}

func Purchase(productID string, amount int) (*Result, error) {
	stock, err := GetStock(productID)
	if err != nil {
		return nil, fmt.Errorf("在庫取得失敗: %w", err) // 技術的例外
	}
	if amount > stock {
		return &Result{OutOfStock: NewOpt(OutOfStockErr{
			Requested: amount,
			Available: stock,
		})}, nil // ビジネス例外
	}
	remaining := stock - amount
	err = SaveStock(productID, remaining)
	if err != nil {
		return nil, fmt.Errorf("在庫保存失敗: %w", err) // 技術的例外
	}
	return &Result{Remaining: NewOpt(remaining)}, nil // 成功
}

type Result struct {
	Remaining  Option[int]
	OutOfStock Option[OutOfStockErr]
}

type OutOfStockErr struct {
	Requested int
	Available int
}
