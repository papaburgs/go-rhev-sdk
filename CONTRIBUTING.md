## Style Notes

For the connection member functions, trying different style of error
handling. Not sure if I like it or not.

Instead of the standard pattern of checking and returing errors all the time,
I am adding an error variable to the struct that can store an error.
Then all the functions will check that error, if it gets set, then we can 
short curcuit out of the functions. Any external functions should still follow
the normal pattern, but internally I think this could streamline this object.

