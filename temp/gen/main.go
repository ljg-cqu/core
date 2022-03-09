// https://github.com/wzshiming/gen
package main

// ItemService #path:"/item/"#
type ItemService struct{}

// Create a Item #route:"POST /"#
func (s *ItemService) Create(item *Item) (err error) {}

// Update the Item #route:"PUT /{item_id}"#
func (s *ItemService) Update(itemID int /* #name:"item_id"# */, item *Item) (err error) {}

// Delete the Item #route:"DELETE /{item_id}"#
func (s *ItemService) Delete(itemID int /* #name:"item_id"# */) (err error) {}

// Get the Item #route:"GET /{item_id}"#
func (s *ItemService) Get(itemID int /* #name:"item_id"# */) (item *ItemWithID, err error) {}

// List of the Item #route:"GET /"#
func (s *ItemService) List(offset, limit int) (items []*ItemWithID, err error) {}
