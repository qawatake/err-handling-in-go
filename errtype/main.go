package main

import "fmt"

func main() {
	remaining, err := Purchase("P-123", 5)
	if err != nil {
		// ビジネス例外はメッセージを伝えて正常終了
		if GetType(err) == OutOfStockErr {
			fmt.Println(err)
			return
		} else {
			panic(err)
		}
	}
	fmt.Printf("残り%d個\n", remaining)
}

// NOTE: 在庫量がamountよりも少ない場合はOutOfStockErrが発生する。
func Purchase(productID string, amount int) (int, error) {
	stock, err := GetStock(productID)
	if err != nil {
		return 0, err // 技術的例外
	}
	if amount > stock {
		return 0, NewErrorf(OutOfStockErr, "購入できません。在庫は残り%d個ですが、%d個必要です。", stock, amount) // ビジネス例外
	}
	remaining := stock - amount
	err = SaveStock(productID, remaining)
	if err != nil {
		return 0, err // 技術的例外
	}
	return remaining, nil // 成功
}

var OutOfStockErr ErrorType = "OutOfStockErr"
