package main

import (
	"aether/nbt"
	"aether/server"
	"compress/gzip"
	"fmt"
	"net"
	"os"
)

func main() {
	f, _ := os.Open("bigtest.nbt")
	gzip, _ := gzip.NewReader(f)
	var value struct {
		N struct {
			Egg struct {
				Name  string  `nbt:"name"`
				Value float32 `nbt:"value"`
			} `nbt:"egg"`
			Ham struct {
				Name  string  `nbt:"name"`
				Value float32 `nbt:"value"`
			} `nbt:"ham"`
		} `nbt:"nested compound test"`
		IntTest    int32   `nbt:"intTest"`
		ByteTest   int8    `nbt:"byteTest"`
		StringTest string  `nbt:"stringTest"`
		LongList   []int64 `nbt:"listTest (long)"`
		DoubleTest float64 `nbt:"doubleTest"`
		FloatTest  float32 `nbt:"floatTest"`
		LongTest   int64   `nbt:"longTest"`
		CompList   []struct {
			CreatedOn int64  `nbt:"created-on"`
			Name      string `nbt:"name"`
		} `nbt:"listTest (compound)"`
		ByteArrayTest []int8 `nbt:"byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))"`
		ShortTest     int16  `nbt:"shortTest"`
	}
	d := nbt.NewDecoder(gzip)
	fmt.Println(d.Decode(&value))
	fmt.Println(value.DoubleTest)
	f.Close()
	gzip.Close()
	cfg := server.ServerConfig{
		IP:                   net.IPv4(127, 0, 0, 1),
		Port:                 25565,
		TPS:                  20,
		CompressionThreshold: -1,
	}
	srv, _ := cfg.New()
	srv.Start()
}
