package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	MiningDifficulty = 3
	MiningReward     = 1.0
	MiningSender     = "THE BLOCKCHAIN"
)

type Transaction struct {
	SenderAddress  string
	ReceiptAddress string
	Value          float64
}

func (t Transaction) Print() {
	fmt.Printf("senderAddress: %s\n", t.SenderAddress)
	fmt.Printf("receiptAddress: %s\n", t.ReceiptAddress)
	fmt.Printf("value: %f\n", t.Value)
}

type Block struct {
	Nonce        int
	PreviousHash [32]byte
	Timestamp    int64
	Transactions []*Transaction
}

type Blockchain struct {
	blockChainAddress string
	blocks            []*Block
	transactionPool   []*Transaction
}

func NewTransaction(senderAddress, receiptAddress string, value float64) *Transaction {
	return &Transaction{
		SenderAddress:  senderAddress,
		ReceiptAddress: receiptAddress,
		Value:          value,
	}
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		Nonce:        nonce,
		PreviousHash: previousHash,
		Timestamp:    time.Now().UnixNano(),
		Transactions: transactions,
	}
}

func (block *Block) Serialize() ([]byte, error) {
	return json.Marshal(block)
}

func (block *Block) Hash() [32]byte {
	body, _ := block.Serialize()
	return sha256.Sum256(body)
}

func (block *Block) Print() {
	fmt.Println("Block nonce:", block.Nonce)
	fmt.Println(fmt.Sprintf("Block hash: %x", block.Hash()))
	fmt.Println(fmt.Sprintf("Block previousHash: %x", block.PreviousHash))
	if len(block.Transactions) == 0 {
		return
	}

	fmt.Println("=======================Transactions ==========================")
	for _, t := range block.Transactions {
		t.Print()
		fmt.Println("___________________________")
	}
}

func NewBlockChain(blockChainAddress string) *Blockchain {
	blockChain := &Blockchain{
		blocks:            []*Block{new(Block)},
		transactionPool:   []*Transaction{},
		blockChainAddress: blockChainAddress,
	}

	return blockChain
}

func (bc *Blockchain) CreateBlock(nonce int) *Block {
	block := NewBlock(nonce, bc.LastBlock().Hash(), bc.transactionPool)
	bc.blocks = append(bc.blocks, block)
	bc.transactionPool = []*Transaction{}

	return block
}

func (bc *Blockchain) AddTransactions(senderAddress string, receiptAddress string, value float64) {
	bc.transactionPool = append(bc.transactionPool, NewTransaction(senderAddress, receiptAddress, value))
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *Blockchain) Print() {
	for i, b := range bc.blocks {
		fmt.Println("--------------------------- chain ", i, "-----------------------------")
		b.Print()
	}
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(t.SenderAddress, t.ReceiptAddress, t.Value))
	}

	return transactions
}

func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{
		Nonce:        nonce,
		PreviousHash: previousHash,
		Transactions: transactions,
	}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())

	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	trxs := bc.CopyTransactionPool()
	prevHash := bc.LastBlock().Hash()
	nonce := 0

	for !bc.ValidProof(nonce, prevHash, trxs, MiningDifficulty) {
		nonce++
	}

	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransactions(MiningSender, bc.blockChainAddress, MiningReward)
	nonce := bc.ProofOfWork()
	bc.CreateBlock(nonce)
	return true
}

func (bc *Blockchain) CalculateTotalAmount(blockChainAddress string) float64 {
	var totalAmount float64

	for _, block := range bc.blocks {
		for _, transaction := range block.Transactions {
			value := transaction.Value
			if transaction.ReceiptAddress == blockChainAddress {
				totalAmount += value
			}

			if transaction.SenderAddress == blockChainAddress {
				totalAmount -= value
			}
		}
	}

	return totalAmount
}
