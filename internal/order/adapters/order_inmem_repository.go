package adapters

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	domain "github.com/sjxiang/oms/order/domain/order"
)


type MemoryOrderRepository struct {
	lock  *sync.RWMutex
	store []*domain.Order
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	
	s := make([]*domain.Order, 0)
	
	s = append(s, &domain.Order{
		ID:          "1",
		CustomerID:  "1",
		Status:      "pending",
		PaymentLink: "",
		Items:       nil,
	})

	return &MemoryOrderRepository{
		lock:  &sync.RWMutex{},
		store: s,
	}
}

func (m *MemoryOrderRepository) Create(_ context.Context, order *domain.Order) (*domain.Order, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	newOrder := &domain.Order{
		ID:          strconv.FormatInt(time.Now().Unix(), 10),
		CustomerID:  order.CustomerID,
		Status:      order.Status,
		Items:       order.Items,
		PaymentLink: order.PaymentLink,
	}
	m.store = append(m.store, newOrder)

	logrus.WithFields(logrus.Fields{
		"input_order":        order,
		"store_after_create": m.store,
	}).Info("memory_order_repo_create")

	return newOrder, nil
}

func (m *MemoryOrderRepository) Get(_ context.Context, id, customerID string) (*domain.Order, error) {
	m.lock.RLock()	
	defer m.lock.RUnlock()

	for _, o := range m.store {
		if o.ID == id && o.CustomerID == customerID {
			logrus.Infof("memory_order_repo_get, found, id=%s, customer_id=%s, res=%+v", id, customerID, *o)
			return o, nil
		}
	}

	return nil, &domain.NotFoundError{OrderID: id}
}

func (m *MemoryOrderRepository)  Update(
	ctx context.Context, 
	order *domain.Order, 
	updateFn func(context.Context, *domain.Order) (*domain.Order, error),
) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	found := false
	for i, o := range m.store {
		
		if o.ID == order.ID && o.CustomerID == order.CustomerID {
			found = true		
			updatedOrder, err := updateFn(ctx, o)
			if err!= nil {
				return err
			}
			m.store[i] = updatedOrder
		}
	}

	if !found {
		return &domain.NotFoundError{OrderID: order.ID}
	}

	return nil
}


