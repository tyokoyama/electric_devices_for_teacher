package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	f, err := os.Create("chokuryu_practice.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		switch t := rand.Int31() % 2; t {
		case 0:
			var v int32
			var Ra float64
			var Ia int32
			for float64(v) <= (float64(Ia) * Ra) {
				v = rand.Int31() % 301
				for v < 100 {
					v = rand.Int31() %301
				}
				Ra = rand.Float64()
				for Ra == 0 {
					Ra = rand.Float64()
				}
				Ia = rand.Int31() % 500
				for Ia == 0 {
					Ia = rand.Int31() %500
				}
			}

			fmt.Fprintf(f, "%d. 誘導起電力%d[V]の分巻発電機がある。電機子巻線の抵抗が%.2f[Ω]、電機子電流が%d[A]の時、端子電圧[V]を求めなさい。ただし、電機子反作用およびブラシによる電圧降下は無視する。\n", (i + 1), v, Ra,Ia)
		case 1:
			var v int32
			var Ra float64
			var Ia int32
			var ea, eb, ef int32
			for float64(v) <= ((float64(Ia) * Ra) + float64(ea) + float64(eb) + float64(ef)) {
				v = rand.Int31() % 301
				for v < 100 {
					v = rand.Int31() %301
				}
				Ra = rand.Float64()
				for Ra == 0 {
					Ra = rand.Float64()
				}
				Ia = rand.Int31() % 500
				for Ia == 0 {
					Ia = rand.Int31() %500
				}
				ea = rand.Int31() % 10
				for ea == 0 {
					ea = rand.Int31() % 10
				}
				eb = rand.Int31() % 10
				for eb == 0 {
					eb = rand.Int31() % 10
				}
				ef = rand.Int31() % 10
				for ef == 0 {
					ef = rand.Int31() % 10
				}
			}

			fmt.Fprintf(f, "%d. 誘導起電力%d[V]の分巻発電機がある。電機子巻線の抵抗が%.2f[Ω]、電機子電流が%d[A]の時、端子電圧[V]を求めなさい。ただし、電機子反作用による電圧降下を%d[V]、ブラシによる電圧降下を%d[V]、界磁電流Ifの減少による電圧降下を%d[V]とする。\n", (i + 1), v, Ra, Ia, ea, eb, ef)

		}
	}
}