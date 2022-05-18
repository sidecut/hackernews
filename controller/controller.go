package controller

import (
	context "context"
)

type Controller struct {
	// Dependencies...
}

// Story struct
type Story struct {
	// Fields...
}

// Index of stories
// GET /
func (c *Controller) Index(ctx context.Context) (stories []*Story, err error) {
	return []*Story{}, nil
}

// Show story
// GET /:id
func (c *Controller) Show(ctx context.Context, id int) (story *Story, err error) {
	return &Story{}, nil
}
