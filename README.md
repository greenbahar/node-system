# node-system
node system - core blockchian

Every node that wants to join the blockchain have to start from the latest block height and download the rest of the blocks from other nodes. So I created separate goroutins to request specific number of blocks in each request from nodes and send them into a channel to be processed by another goroutine to decide add the received blocks into the blockchain history of the new node. Nodes talk to each other via gRPC. 
