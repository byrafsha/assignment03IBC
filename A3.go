package assignment03IBC

import (
	"crypto/sha256"
	"fmt"
	"encoding/gob"
	"net"
	a2 "github.com/rafshamazhar/assignment02IBC"
)
var Quorum int
var chainHead *a2.Block
storeMap:=make(map[string]net.Conn)
channel1 := make(chan net.Conn)
channel2:= make(chan string)
	//[]byte, 4096)

func handleConnection(c net.Conn, node string, listeningAddress string){
	if node=="satoshi" {
		//store connection&address
		//receive listening port on c first
		storemap[ReadString(c)]=c

		//mine new block
		chainHead = a2.InsertBlock("", "", "Satoshi", 0, chainHead)

		//print address
		log.Println("Satoshi: Client connected at ". c.RemoteAddr())

	} else if node == "others" {
		//others: satoshi->net.conn 
		//which port node is listening to, msg recieve and print
		//read and print msg
		log.Println("Others: Client connected at ". c.RemoteAddr())
	}
}

func StartListening(listeningAddress string, node string) {
	ln, err:= net.Listen("tcp", listeningAddress)
	if err!=nil{
		log.Fatal(err)
	}
	for {
		conn, err:=ln.Accept()
		if err!=nil{
			log.Println(err)
			continue
		}
		go handleConnection(conn, node, listeningAddress)  //storage if satoshi, only print if others
	}
}


//WaitForQuorum()
func WaitForQuorum() {
	if len(storeMap<Quorum){
		x,y:=<-channel2,<-channel2
	}
}


func SendChainandConnInfo() {

	//blockchain using gob
	for _, val in range(storeMap) {
		blockchainEnc := gob.NewEncoder(&val) //loop through all c conn
		err:= blockchainEnc.Encode(chainHead)
		if err!=nil {
			log.Fatal("encode error:", err)
		}
	}

	//connection topology
}


func ReceiveChain(connection net.Conn) *Block {
	blockchainDec := gob.NewDecoder(&connection)
	var mychain *Block
	err := blockchainDec.Decode(&mychain)
	if err!=nil{
		log.Fatal("decode error", err)
	}
	return mychain
}


func WriteString(connection net.Conn, listeningAddress string) {
	channel1 <- connection
	channel2 <- listeningAddress
}


func ReadString(connection net.Conn) string {
	x := <-channel1
	y := <-channel1
	if x==connection{
		return y
	}
}
