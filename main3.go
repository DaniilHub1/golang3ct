package main

func main() {
	kan := make(chan int)
	
	kan <- 1
}

// решение дедлока ну вдруг интересно 
//  func main() {
//     kan := make(chan int)

//     go func() {
//         kan <- 1
//     }()

//     val := <-kan
//     println(val)
// }
