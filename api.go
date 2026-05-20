package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.1" )

func maF4bQODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SgK6DeP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BtZJbnjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bdNdvSXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zt83lsZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nm8GntXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jFcmqAN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SGHMdXGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 13yA4M0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0ysyHBqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W1jmF6ZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6OsOBkLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MjwEOcG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func okstlq6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dsWxPYjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ZbLLfdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uP5CkwJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ktrC3skEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3leEUhNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aMFGkI0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oe3ZAAFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pKtUcWluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 87J1zHGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X40eA5q0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uFxhNuhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LYt3l6TyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YrfH3CmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3DlMICCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hl3Ml5stWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1gepZJlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bgLnivf1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JdNIHgmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F8QOwR2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kG1yeDOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vie9uuQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UzIVWJMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gmZquY6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wNRttDFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vaL0ODN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func blY0ipvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q2A8j6f3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IViciowBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uEb3AyeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZCxgvg9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 480zbUNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UosOOZNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z9ur6amJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zTvaS2nbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1QKBBq0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ElOeFvzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YXFHTJwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yTsj0CrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DQAM4RJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 21ekzaKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GkRvaXpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K4tff7EGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KQe104cQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N9R837CJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ReASPcxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vSV1xj4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6xpaVIv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V5i74TkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WIz3QBLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i3PepWxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DXUSEfsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MSUoUlAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jiD03HG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PrrFwOLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EshyuJ5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VV9KajAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func boPowwqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rKJR7gy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wt2Vc8w1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dfWWZA8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FRXwPBAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func joCqVq5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mkSr9UrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hEDk3HHlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 76ZI89C6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8MQuORQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NagWMPVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cv05t2s1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HsPWhkpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EXSVMwdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FwZPhnO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fomQP3X2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pNrBD3SWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CsHS9JsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oeJ21TV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func imBM5pXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 76eOw6wsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DanitMGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZHmcXk9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UCt0vSOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zgtdoX3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rkNxZaJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dvQtR0XJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oVIKRJDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v7t0SZFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oRXN89uvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uE6lIvbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZTsHvfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AxmbcsOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fh6OOIKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZDzn1D1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sno06TpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uxr7JzYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X9PhThlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OeGDYNSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AGpoYymOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8X5xoBLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ydDMZ4noWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lFNgmt6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QmjKU5BiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xaltREZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5QJxQSMdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hPisim9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C3shx2XzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iFGFECpSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func puPRanUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3kmpx1aeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BBGlDBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xfeqekuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ICEx8ukHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PnqsOrKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G9KMSDtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vTxr9gT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qqRl19gJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K6YFeawCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HETHDBrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r1s47ooAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QHZSHRw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iHZ6G0otWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RwW6RoXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KVa4JelqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ck4lTr5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HfyAPfhEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0FAU2E7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1UCytBtJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HPdtVcHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3mOmDy6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6B3h7hfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 00wTPAOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FPZZoAM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YY5VOq2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e8T0ktIOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Y1DTrEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YuTZo5uTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GbB7wTnaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ByyzHW1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9BlG18ABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EqttJWJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TJBXrPbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHC4uDDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZWdII9eSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HMnq9OSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tAQseWOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a8EUQoHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dkHVyp2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mtaS8Hb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func phh0aBbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ci5yAKEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 77yK4JTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nWxvOUsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cGSmSsBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3KzoN5ILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tS3Ay6LWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zpPnchT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OZFpyg0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hsR3XlOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EQ5gQQ8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3RlVTusIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yajSVzCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8BSiTVoUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHe24h7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mRaqD5GCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IowNweNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V6EakCddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sdWJKdKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rm6aJyEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tG93LC3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0SmyshCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BBGNGaHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lnGWUKTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pZmPLEP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dgJobJa1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z9sZTH9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6uMkFgu7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func psBaSTGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HXBThDHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oatxEZMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NwJ1M5rDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XNRtlF0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mcrhJnqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uCx3CBY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HpbK3cDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QM6Z7ajCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 26r8MbVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JfIXHV03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kah2zGhTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DGtoou8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YDg71Hw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M2FjMqsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UMIAyKQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

