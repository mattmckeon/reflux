package main

import (
	"fmt"
	"io"
)

// Transaction represents a reversible set of operations.
type Transaction struct {
	RollbackLog []Operation
}

// NewTransaction creates and initializes a new Transaction.
func NewTransaction() *Transaction {
	return &Transaction{
		RollbackLog: []Operation{},
	}
}

// AddToRollbackLog adds the given operation to this transaction's rollback log.
// If the transaction is rolled back, operations will be applied in LIFO order.
func (t *Transaction) AddToRollbackLog(op Operation) {
	t.RollbackLog = append([]Operation{op}, t.RollbackLog...)
}

// Rollback rolls back this transaction.
func (t *Transaction) Rollback(data *RefluxData) {
	for _, op := range t.RollbackLog {
		op.Apply(data)
	}
}

// RefluxDb is a transaction layer atop a simple key-value datastore.
type RefluxDb struct {
	Data         *RefluxData
	Transactions []*Transaction
}

// NewRefluxDb creates a new instance of RefluxDb with a clean datastore and
// transaction log.
func NewRefluxDb() *RefluxDb {
	return &RefluxDb{
		Data:         NewRefluxData(),
		Transactions: []*Transaction{},
	}
}

// DoGet executes a GET operation, writing output to the given Writer.
func (r *RefluxDb) DoGet(key string, writer io.Writer) {
	(&Get{Key: key, Output: writer}).Apply(r.Data)
}

// DoSet executes a SET operation.
func (r *RefluxDb) DoSet(key, value string) {
	inverse := (&Set{Key: key, Value: value}).Apply(r.Data)
	for i := range r.Transactions {
		r.Transactions[i].AddToRollbackLog(inverse)
	}
}

// DoDelete executes a DELETE operation.
func (r *RefluxDb) DoDelete(key string) {
	inverse := (&Delete{Key: key}).Apply(r.Data)
	for i := range r.Transactions {
		r.Transactions[i].AddToRollbackLog(inverse)
	}
}

// DoCount executes a COUNT operation, writing output to the given Writer.
func (r *RefluxDb) DoCount(value string, writer io.Writer) {
	(&Count{Value: value, Output: writer}).Apply(r.Data)
}

// DoBegin executes a BEGIN operation.
func (r *RefluxDb) DoBegin() {
	r.Transactions = append(r.Transactions, NewTransaction())
}

// DoBegin executes a COMMIT operation.
func (r *RefluxDb) DoCommit() {
	r.Transactions = r.Transactions[:len(r.Transactions)-1]
}

// DoRollback executes a ROLLBACK operation, writing output to the given Writer.
func (r *RefluxDb) DoRollback(writer io.Writer) {
	if len(r.Transactions) < 1 {
		fmt.Fprintln(writer, "TRANSACTION NOT FOUND")
	} else {
		r.Transactions[len(r.Transactions)-1].Rollback(r.Data)
		r.Transactions = r.Transactions[:len(r.Transactions)-1]
	}
}
