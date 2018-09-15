package main

import ()

type Reflux struct {
	Data  map[string]string
	Counts map[string]int
}

func NewReflux() *Reflux {
	return &Reflux{
		Data: make(map[string]string),
		Counts: make(map[string]int),
	}
}

func (r *Reflux) Set(key, value string) {
	r.Delete(key)
	r.Data[key] = value
	if count, hasCount := r.Counts[value]; hasCount {
		r.Counts[value] = count + 1
	} else {
		r.Counts[value] = 1
	}
}

func (r *Reflux) Has(key string) bool {
	_, exists := r.Data[key]
	return exists
}

func (r *Reflux) Get(key string) string {
	return r.Data[key]
}

func (r *Reflux) Delete(key string) {
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

func (r *Reflux) Count(value string) int {
	return r.Counts[value]
}
