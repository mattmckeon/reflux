package main

import ()

// RefluxData is the datastore layer of the database.
type RefluxData struct {
	Data   map[string]string
	Counts map[string]int
}

// NewRefluxData constructs a new, empty instance of the datastore.
func NewRefluxData() *RefluxData {
	return &RefluxData{
		Data:   make(map[string]string),
		Counts: make(map[string]int),
	}
}

// Set sets the given key to the given value in the datastore.
func (r *RefluxData) Set(key, value string) {
	r.Delete(key)
	r.Data[key] = value
	if count, hasCount := r.Counts[value]; hasCount {
		r.Counts[value] = count + 1
	} else {
		r.Counts[value] = 1
	}
}

// Has returns whether the given key is mapped in the datastore.
func (r *RefluxData) Has(key string) bool {
	_, exists := r.Data[key]
	return exists
}

// Get returns the current value for the given key.
func (r *RefluxData) Get(key string) string {
	return r.Data[key]
}

// Delete removes the key and its value from the datastore.
func (r *RefluxData) Delete(key string) {
	if currentValue, hasValue := r.Data[key]; hasValue {
		if count, hasCount := r.Counts[currentValue]; hasCount {
			if count == 1 {
				delete(r.Counts, currentValue)
			} else {
				r.Counts[currentValue] = count - 1
			}
		}
		delete(r.Data, key)
	}
}

// Count returns the number of times the given value occurs in the datastore.
func (r *RefluxData) Count(value string) int {
	return r.Counts[value]
}
