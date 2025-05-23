package store

import (
	"context"
	"todolist/pkg/structs"
)

type Store interface {
	Update(action func(tx Txn) error) error
}

type Txn interface {
	Add(ctx context.Context, item *structs.TodoItem) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, e *structs.TodoItem) error
	Get(ctx context.Context, id string, item *structs.TodoItem) error
	List(ctx context.Context, items *structs.TodoItemList) error
	ReorderItem(ctx context.Context, id string, newOrder int) error
	DbTx() interface{}
}
