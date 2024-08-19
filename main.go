package main

import (
	"fmt"

	"github.com/patriciabonaldy/blockchain/blockchain"
)

func main() {
	blockChain := blockchain.NewBlockChain("myBlockChainAddress")

	blockChain.AddTransactions("A", "B", 1.0)
	blockChain.Mining()

	blockChain.AddTransactions("C", "D", 2.0)
	blockChain.AddTransactions("E", "F", 3.0)
	blockChain.Mining()

	blockChain.Print()

	fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", blockChain.CalculateTotalAmount("D"))
	fmt.Printf("A %.1f\n", blockChain.CalculateTotalAmount("A"))
	fmt.Printf("A %.1f\n", blockChain.CalculateTotalAmount("A"))
	//pageCount(4, 4)
}

func pageCount(n int32, p int32) int32 {
	mid := n / 2
	pag := p / 2

	fmt.Println("mid--", mid, "pag--", pag)
	if p <= mid {
		return pag
	}

	return mid - pag
}
