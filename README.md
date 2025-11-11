# maybe

This project provides a generic container that boxes a value Some or nil type None.
A typical way to do this in go is to a pointer to a value and use the default pointer value, nil, for None.
I personally find this a little clunky and hard to look at as it usually uses a small function Ptr(val T) *T to wrap each of the values.
