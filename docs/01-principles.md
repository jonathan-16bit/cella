# Durability and Atomicity

Writing to a file does not necessarily always mean the data is safely saved. 
The computer may hold the change in RAM for a while, so a power loss could erase it.  

Durability: Once the DB confirms data is saved, the change should survive a crash (because its in storage).  

Atomicity: If a crash happens halfway through a change, the database should end up with either 
the whole change or none of it, not half a change.  

`fsync`: Flushes pending changes from OS memory to storage, then waits for storage to report that its done. 
However, it can't make multiple writes be done as one atomic action.   
