package main

import (
	"bytes"
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
	"log"
	"os"
)

const lenUser = 4

type User struct {
	Name    string
	Age     int
	Address string
	House   int
}

func main() {
	u := &User{
		Name:    "John",
		Age:     42,
		Address: "Bryansk",
		House:   255,
	}

	buf := &bytes.Buffer{}

	err := msgpack.NewEncoder(buf).Encode(u)
	if err != nil {
		log.Printf("NewEncoder %v", err)
		os.Exit(1)
	}
	fmt.Println("Encode:")
	fmt.Printf("% x\n", buf)
	fmt.Printf("%s\n", buf)

	uDec := &User{}
	err = msgpack.NewDecoder(buf).Decode(uDec)
	if err != nil {
		log.Printf("NewDecoder %v", err)
		os.Exit(1)
	}

	fmt.Printf("\nDecode:\n")
	fmt.Println(uDec.Name, uDec.Age, uDec.Address, uDec.House)

}

//========================================================================
// Запусти программу. Раскомментируй функции и снова запусти программу...
//========================================================================
//
//func (u *User) DecodeMsgpack(d *msgpack.Decoder) error {
//	var err error
//	var fieldsCount int
//
//	if fieldsCount, err = d.DecodeArrayLen(); err != nil {
//		return err
//	}
//	if fieldsCount < lenUser {
//		return fmt.Errorf("binder session has wrong length %d, expect min %d", fieldsCount, lenUser)
//	}
//	if u.Name, err = d.DecodeString(); err != nil {
//		return err
//	}
//	if u.Age, err = d.DecodeInt(); err != nil {
//		return err
//	}
//	if u.Address, err = d.DecodeString(); err != nil {
//		return err
//	}
//	if u.House, err = d.DecodeInt(); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (u *User) EncodeMsgpack(e *msgpack.Encoder) error {
//	if err := e.EncodeArrayLen(lenUser); err != nil {
//		return err
//	}
//	if err := e.EncodeString(u.Name); err != nil {
//		return err
//	}
//	if err := e.EncodeInt(u.Age); err != nil {
//		return err
//	}
//	if err := e.EncodeString(u.Address); err != nil {
//		return err
//	}
//	if err := e.EncodeInt(u.House); err != nil {
//		return err
//	}
//	return nil
//}
