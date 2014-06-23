package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	f, err := os.Create("yuudou_practice.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		var p int32
		var z int32
		var n int32
		var e int32

		p = ((rand.Int31() % 4) + 1) * 2

		for z == 0 {
			z = rand.Int31() % 1000
		}

		for n == 0 {
			n = rand.Int31() % 3000
		}

		for e == 0 {
			e = rand.Int31() % 500
		}

		maki := rand.Int31() % 2
		if maki == 0 {
			fmt.Fprintf(f, "磁極数が%d、波巻、電機子全導線数が%dの直流発電機を%d[r/min]の回転速度で運転した時、発電機の起電力は%d［V］であった。この時の1極当たりの磁束Φ[Wb]を求めなさい。\n", p, z, n, e)
		} else {
			fmt.Fprintf(f, "磁極数が%d、重ね巻、電機子全導線数が%dの直流発電機を%d[r/min]の回転速度で運転した時、発電機の起電力は%d［V］であった。この時の1極当たりの磁束Φ[Wb]を求めなさい。\n", p, z, n, e)
		}
	}
}