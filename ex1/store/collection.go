package store

import (
	"errors"
)

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

var documents = map[string]Document{}

var ErrDocumentNotFound = errors.New("document not found")
var ErrInvalidDocumentInput = errors.New("invalid document input or output is nil")

type DocumentField struct {
	Type  DocumentFieldType
	Value interface{}
}

type Document struct {
	Fields map[string]DocumentField
}

type Collection struct {
	documents  map[string]Document
	primaryKey string
}

type MyCollection interface {
	Create(doc Document)
	Get(key string) (*Document, bool)
	Delete(key string) bool
	List() []Document
}

func NewCollection() *Collection {
	return &Collection{
		documents: map[string]Document{},
	}
}

func (c *Collection) Create(doc Document) {
	keyField, ok := doc.Fields["key"]
	if !ok || keyField.Type != DocumentFieldTypeString {
		return
	}
	keyStr, ok := keyField.Value.(string)
	if !ok {
		return
	}
	c.documents[keyStr] = doc
}

func (c *Collection) Get(key string) (*Document, bool) {
	doc, found := c.documents[key]
	if !found {
		return nil, false
	}
	return &doc, true
}

func (c *Collection) Delete(key string) bool {
	_, found := c.documents[key]
	if found {
		delete(c.documents, key)
		return true
	}
	return false
}

func (c *Collection) List() []Document {
	result := make([]Document, 0, len(c.documents))
	for _, doc := range c.documents {
		result = append(result, doc)
	}
	return result
}
