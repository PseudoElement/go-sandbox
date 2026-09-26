package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	gopatterns "github.com/pseudoelement/go-sandbox/go-patterns"
)

var (
	ErrNotEnoughBalance      = errors.New("not enough balance")
	ErrUserAlreadyProcessing = errors.New("user already processing")
)

var txs = []Transaction{
	{ID: UserID(strconv.Itoa(int(rand.Uint64()))), UserID: "1", Amount: 100},
	{ID: UserID(strconv.Itoa(int(rand.Uint64()))), UserID: "2", Amount: 200},
	{ID: UserID(strconv.Itoa(int(rand.Uint64()))), UserID: "3", Amount: 300},
	{ID: UserID(strconv.Itoa(int(rand.Uint64()))), UserID: "4", Amount: 400},
	{ID: UserID(strconv.Itoa(int(rand.Uint64()))), UserID: "1", Amount: 500},
	{ID: UserID(strconv.Itoa(int(rand.Uint64()))), UserID: "2", Amount: 600},
	{ID: UserID(strconv.Itoa(int(rand.Uint64()))), UserID: "3", Amount: 700},
	{ID: UserID(strconv.Itoa(int(rand.Uint64()))), UserID: "4", Amount: 800},
	{ID: UserID(strconv.Itoa(int(rand.Uint64()))), UserID: "1", Amount: 800},
	{ID: UserID(strconv.Itoa(int(rand.Uint64()))), UserID: "2", Amount: 800},
	{ID: UserID(strconv.Itoa(int(rand.Uint64()))), UserID: "3", Amount: 800},
	{ID: UserID(strconv.Itoa(int(rand.Uint64()))), UserID: "4", Amount: 800},
}

type DB struct{}

func (db *DB) Query(string) {}

func main_processor() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	balances := map[UserID]int{
		"1": 100000,
		"2": 100000,
		"3": 200000,
		"4": 300000,
	}
	inbox := make(chan Transaction)
	processor := NewProcessor(4, balances)
	// ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	ctx2, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			for _, tx := range txs {
				inbox <- tx
			}
			time.Sleep(2 * time.Second)
		}
	}()

Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		case tx := <-inbox:
			go func() {
				if err := processor.Submit(ctx2, tx); err != nil {
					fmt.Printf("[%s] err: %v\n", tx.UserID, err)
				} else {
					fmt.Printf("[%s] submitted\n", tx.UserID)
				}
			}()
		}
	}

	<-ctx.Done()

	log.Println("Server is shutting down...")
	wg := sync.WaitGroup{}
	for range 100 {
		wg.Add(1)
		go func() {
			processor.Shutdown()
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("Server exited cleanly!")
}

type UserID string

type Transaction struct {
	ID     UserID
	UserID UserID
	Amount int64
}

type PendingUser struct {
	operation func() error
	ch        chan struct{}
}

type Processor struct {
	shutdownInProcess atomic.Bool
	balances          map[UserID]int
	mu                sync.Mutex
	pendingUsers      map[UserID]PendingUser
	db                *DB
	sem               *gopatterns.Semaphore
}

func NewProcessor(workers int, balances map[UserID]int) *Processor {
	return &Processor{
		shutdownInProcess: atomic.Bool{},
		db:                &DB{},
		balances:          balances,
		pendingUsers:      make(map[UserID]PendingUser),
		sem:               gopatterns.NewSemaphore(context.TODO(), workers),
	}
}

func (p *Processor) Submit(ctx context.Context, tx Transaction) error {
	err := p.sem.Acquire()
	if err != nil {
		return err
	}
	defer p.sem.Release()

	p.mu.Lock()
	user, ok := p.pendingUsers[UserID(tx.UserID)]
	if !ok {
		user = PendingUser{
			ch: make(chan struct{}, 1),
		}
		p.pendingUsers[UserID(tx.UserID)] = user
	}
	user.operation = func() error { return p._process(tx) }
	p.mu.Unlock()

	err = p._wait(ctx, user, func() error {
		return p._process(tx)
	})

	return err
}

func (p *Processor) _process(tx Transaction) error {
	p.mu.Lock()
	userBalance, ok := p.balances[UserID(tx.UserID)]
	if !ok {
		p.balances[tx.UserID] = 0
		userBalance = 0
	}
	p.mu.Unlock()

	if userBalance < int(tx.Amount) {
		return ErrNotEnoughBalance
	}

	time.Sleep(time.Duration(1) * time.Second)

	p.mu.Lock()
	newBalance := userBalance - int(tx.Amount)
	p.balances[tx.UserID] = newBalance
	p.mu.Unlock()

	return nil
}

func (p *Processor) _wait(ctx context.Context, user PendingUser, process func() error) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case user.ch <- struct{}{}:
		err := process()
		<-user.ch
		return err
	}
}

func (p *Processor) Shutdown() {
	if !p.shutdownInProcess.CompareAndSwap(false, true) {
		return
	}
	println("Processor Shutdown started...")
	p.mu.Lock()
	defer p.mu.Unlock()
	wg := sync.WaitGroup{}
	for _, user := range p.pendingUsers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(2 * time.Second)
			if err := user.operation(); err != nil {
				p.db.Query("INSERT into ops_errors...")
			} else {
				p.db.Query("UPDATE ops...")
			}
		}()
	}
	wg.Wait()
	println("Processor Shutdown completed!")
}
