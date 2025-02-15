package store

import (
	"context"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sqlitedb "todolist/pkg/db"
	"todolist/pkg/structs"
)

var testingT *testing.T

func TestTodoSqlStorage(t *testing.T) {
	testingT = t
	RegisterFailHandler(Fail)

	RunSpecs(t, "store suite")
}

var _ = Describe("SQL Store tests", func() {
	var tododb *sqlx.DB
	var todostore Store
	var ctx context.Context

	Context("When database created", Ordered, func() {

		BeforeAll(func() {
			var err error
			tododb, err = sqlitedb.CreateDb()
			Expect(err).NotTo(HaveOccurred())
			todostore = NewSQLStore(tododb)
			ctx = context.Background()
		})

		AfterAll(func() {
			err := tododb.Close()
			Expect(err).NotTo(HaveOccurred())
		})

		Specify("List returns empty", func() {
			var items structs.TodoItemList

			err := todostore.Update(func(tx Txn) error {
				return tx.List(ctx, &items)
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(items.Items).To(BeEmpty())
			Expect(items.Count).To(Equal(0))
		})

		Context("When todo item created", func() {
			var item structs.TodoItem
			BeforeEach(func() {
				item = structs.TodoItem{Id: "7efc0335-8da6-45f7-a9b6-d4a46ba3044b", Item: "Service motorbike", Item_order: 1}
				err := todostore.Update(func(tx Txn) error {
					return tx.Add(ctx, &item)
				})
				Expect(err).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				item = structs.TodoItem{Id: "7efc0335-8da6-45f7-a9b6-d4a46ba3044b", Item: "Service motorbike", Item_order: 1}
				err := todostore.Update(func(tx Txn) error {
					return tx.Delete(ctx, item.Id)
				})
				Expect(err).NotTo(HaveOccurred())
			})

			Specify("Item is returned from get", func() {
				var gItem structs.TodoItem
				err := todostore.Update(func(tx Txn) error {
					return tx.Get(ctx, item.Id, &gItem)
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(gItem).To(Equal(item))
			})

			Specify("Item is returned from List", func() {
				var items structs.TodoItemList

				err := todostore.Update(func(tx Txn) error {
					return tx.List(ctx, &items)
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(items.Count).To(Equal(1))
				Expect(items.Items).To(ContainElement(item))
			})

			Context("When todo item modified", func() {
				var updatedItem structs.TodoItem
				BeforeEach(func() {
					updatedItem = structs.TodoItem{Id: "7efc0335-8da6-45f7-a9b6-d4a46ba3044b", Item: "Service motorbike and book MOT", Item_order: 1}
					err := todostore.Update(func(tx Txn) error {
						return tx.Update(ctx, &updatedItem)
					})
					Expect(err).NotTo(HaveOccurred())
				})

				Specify("Item is returned from get", func() {
					var gItem structs.TodoItem
					err := todostore.Update(func(tx Txn) error {
						return tx.Get(ctx, item.Id, &gItem)
					})

					Expect(err).NotTo(HaveOccurred())
					Expect(gItem).To(Equal(updatedItem))
					Expect(gItem).NotTo(Equal(item))
				})
			})

			Context("When second todo item created", func() {
				var secondItem structs.TodoItem
				BeforeEach(func() {
					secondItem = structs.TodoItem{Id: "dac2581f-9c76-47aa-877e-6c15ddcfb064", Item: "Book holiday", Item_order: 2}
					err := todostore.Update(func(tx Txn) error {
						return tx.Add(ctx, &secondItem)
					})
					Expect(err).NotTo(HaveOccurred())
				})

				AfterEach(func() {
					err := todostore.Update(func(tx Txn) error {
						return tx.Delete(ctx, secondItem.Id)
					})
					Expect(err).NotTo(HaveOccurred())
				})

				Specify("Item is returned from get", func() {
					var gItem structs.TodoItem
					err := todostore.Update(func(tx Txn) error {
						return tx.Get(ctx, secondItem.Id, &gItem)
					})

					Expect(err).NotTo(HaveOccurred())
					Expect(gItem).To(Equal(secondItem))
				})

				Specify("Item is returned from List", func() {
					var items structs.TodoItemList

					err := todostore.Update(func(tx Txn) error {
						return tx.List(ctx, &items)
					})
					Expect(err).NotTo(HaveOccurred())
					Expect(items.Count).To(Equal(2))
					Expect(items.Items).To(ContainElements(item, secondItem))
				})
			})

			Context("When reorder operation is performed", func() {
				// first I thought I was making some mistake but after the above tests run there is an element that stays in db have considered it an an entry and went forward with the reordering. Hope thats fine!

				var item1, item2, item3 structs.TodoItem

				BeforeEach(func() {
					item1 = structs.TodoItem{Id: "6e57f1df-6713-4ad8-a9a1-3bae1b0b4568", Item: "First item", Item_order: 2}
					item2 = structs.TodoItem{Id: "9c4d5c05-dbb4-4dbf-bfd9-d2a1dfdadbba", Item: "Second item", Item_order: 3}
					item3 = structs.TodoItem{Id: "7dbaa2b1-9fe0-4033-9401-3fdc5b63b590", Item: "Third item", Item_order: 4}

					err := todostore.Update(func(tx Txn) error {
						if err := tx.Add(ctx, &item1); err != nil {
							return err
						}
						if err := tx.Add(ctx, &item2); err != nil {
							return err
						}
						return tx.Add(ctx, &item3)
					})
					Expect(err).NotTo(HaveOccurred())
				})

				AfterEach(func() {
					err := todostore.Update(func(tx Txn) error {
						if err := tx.Delete(ctx, item1.Id); err != nil {
							return err
						}
						if err := tx.Delete(ctx, item2.Id); err != nil {
							return err
						}
						return tx.Delete(ctx, item3.Id)
					})
					Expect(err).NotTo(HaveOccurred())
				})

				Specify("Reordering adjusts all items correctly", func() {
					item3.Item_order = 1
					err := todostore.Update(func(tx Txn) error {
						return tx.ReorderItem(ctx, "7dbaa2b1-9fe0-4033-9401-3fdc5b63b590", 1)
					})
					Expect(err).NotTo(HaveOccurred())

					var items structs.TodoItemList
					err = todostore.Update(func(tx Txn) error {
						return tx.List(ctx, &items)
					})
					Expect(err).NotTo(HaveOccurred())
					Expect(items.Count).To(Equal(4))
					Expect(items.Items).To(ConsistOf(
						structs.TodoItem{Id: "7efc0335-8da6-45f7-a9b6-d4a46ba3044b", Item: "Service motorbike", Item_order: 2},
						structs.TodoItem{Id: "6e57f1df-6713-4ad8-a9a1-3bae1b0b4568", Item: "First item", Item_order: 3},
						structs.TodoItem{Id: "9c4d5c05-dbb4-4dbf-bfd9-d2a1dfdadbba", Item: "Second item", Item_order: 4},
						structs.TodoItem{Id: "7dbaa2b1-9fe0-4033-9401-3fdc5b63b590", Item: "Third item", Item_order: 1},
					))
				})
			})

		})
	})
})
