package main

//func main() {
//	//i := make([]int, 10, 100)
//	//fmt.Println(i[9])
//	//
//	//ints := make(chan int, 100)
//	////通过Context，chan进行通信，golang建议。
//	//cancel, cancelFunc := context.WithCancel(context.TODO())
//	//for i := 0; i < cap(ints); i++ {
//	//	go func(i int) {
//	//		ints <- i
//	//		fmt.Printf("send %v to chan ", i)
//	//		if i == cap(ints)-1 {
//	//			fmt.Println("close chan , invoke cancel function")
//	//			close(ints)
//	//			cancelFunc()
//	//		}
//	//	}(i)
//	//	time.Sleep(time.Second / 10)
//	//}
//	//
//	///**
//	//读取不能知道写入多少
//	//*/
//	//select {
//	//case <-cancel.Done():
//	//	fmt.Println("begin read")
//	//	for i := range ints {
//	//		fmt.Print(i)
//	//	}
//	//}
//}

//func main() {
//	c := exec.Command("cmd", "/C", "ipconfig")
//	//c := exec.Command("cmd", "/C", "ping", "www.baidu.com")
//	r, w, err := os.Pipe()
//	if err != nil {
//		panic(err.Error())
//	}
//
//	c.Stdout = w
//	c.Stderr = w
//	//c.Stdin = os.Stdin
//	go func() {
//		fmt.Println("start read ")
//		reader := bufio.NewReader(r)
//		for {
//			s, err2 := reader.ReadString('\n')
//			fmt.Printf("%s", s)
//			if err2 != nil {
//				print(err.Error())
//				break
//			}
//		}
//		//bytes := make([]byte, 8)
//		//for {
//		//	read, err := r.Read(bytes)
//		//	if err != nil && err == io.EOF {
//		//		break
//		//	}
//		//	if err != nil {
//		//		panic(err.Error())
//		//	}
//		//	fmt.Printf("%s", bytes[0:read])
//		//}
//		fmt.Println("end read ")
//	}()
//	err = c.Start()
//	if err != nil {
//		print(err.Error())
//	}
//	err = c.Wait()
//	if err != nil {
//		panic(err.Error())
//	}
//}

//func main() {
//	command, _, cancelFunc := basecmd.NewCommand()
//	go func() {
//		time.Sleep(time.Second * 10)
//		cancelFunc()
//	}()
//	err := command.Execute()
//	if err != nil {
//		panic(err.Error())
//	}
//}
