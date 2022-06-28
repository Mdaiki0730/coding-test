package main

import (
  "fmt"
  "github.com/ethereum/go-ethereum/p2p"
)

func main() {
  rw1, rw2 := p2p.MsgPipe()
  go func() {
  	p2p.Send(rw1, 8, [][]byte{{0, 0}})
  	p2p.Send(rw1, 5, [][]byte{{1, 1}})
  	rw1.Close()
  }()

  for {
  	msg, err := rw2.ReadMsg()
  	if err != nil {
  		break
  	}
  	var data [][]byte
  	msg.Decode(&data)
  	fmt.Printf("msg: %d, %x\n", msg.Code, data[0])
  }
}
