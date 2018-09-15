# reflux
Not exactly ACID

An in-memory key/value store with transactions and value reference counting.
Created as a coding exercise while learning Go.

- [main.go](main.go): Creates the datastore, reads commands from stdin and applies them to the datastore.
- [reflux.go](reflux.go): The transaction layer of the datastore.
- [operations.go](operations.go): Defines a set of reversible datastore operations.
- [reflux_data.go](reflux_data.go): The key/value store layer.
