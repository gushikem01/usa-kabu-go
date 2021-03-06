// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/gushikem01/usa-kabu-go/server/ent/migrate"

	"github.com/gushikem01/usa-kabu-go/server/ent/stocks"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Stocks is the client for interacting with the Stocks builders.
	Stocks *StocksClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Stocks = NewStocksClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Stocks: NewStocksClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Stocks: NewStocksClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Stocks.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Stocks.Use(hooks...)
}

// StocksClient is a client for the Stocks schema.
type StocksClient struct {
	config
}

// NewStocksClient returns a client for the Stocks from the given config.
func NewStocksClient(c config) *StocksClient {
	return &StocksClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `stocks.Hooks(f(g(h())))`.
func (c *StocksClient) Use(hooks ...Hook) {
	c.hooks.Stocks = append(c.hooks.Stocks, hooks...)
}

// Create returns a create builder for Stocks.
func (c *StocksClient) Create() *StocksCreate {
	mutation := newStocksMutation(c.config, OpCreate)
	return &StocksCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Stocks entities.
func (c *StocksClient) CreateBulk(builders ...*StocksCreate) *StocksCreateBulk {
	return &StocksCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Stocks.
func (c *StocksClient) Update() *StocksUpdate {
	mutation := newStocksMutation(c.config, OpUpdate)
	return &StocksUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StocksClient) UpdateOne(s *Stocks) *StocksUpdateOne {
	mutation := newStocksMutation(c.config, OpUpdateOne, withStocks(s))
	return &StocksUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StocksClient) UpdateOneID(id int) *StocksUpdateOne {
	mutation := newStocksMutation(c.config, OpUpdateOne, withStocksID(id))
	return &StocksUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Stocks.
func (c *StocksClient) Delete() *StocksDelete {
	mutation := newStocksMutation(c.config, OpDelete)
	return &StocksDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StocksClient) DeleteOne(s *Stocks) *StocksDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StocksClient) DeleteOneID(id int) *StocksDeleteOne {
	builder := c.Delete().Where(stocks.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StocksDeleteOne{builder}
}

// Query returns a query builder for Stocks.
func (c *StocksClient) Query() *StocksQuery {
	return &StocksQuery{
		config: c.config,
	}
}

// Get returns a Stocks entity by its id.
func (c *StocksClient) Get(ctx context.Context, id int) (*Stocks, error) {
	return c.Query().Where(stocks.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StocksClient) GetX(ctx context.Context, id int) *Stocks {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStocks queries the stocks edge of a Stocks.
func (c *StocksClient) QueryStocks(s *Stocks) *StocksQuery {
	query := &StocksQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(stocks.Table, stocks.FieldID, id),
			sqlgraph.To(stocks.Table, stocks.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, stocks.StocksTable, stocks.StocksPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StocksClient) Hooks() []Hook {
	return c.hooks.Stocks
}
