package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func openPrimaryAndBackupFile(i int) {
	if i == 0{
	    os.Create("primary.txt")
	    addNumber(i)
	} else{
		 os.Rename("backup.txt", "primary.txt")
	}
	os.Open("primary.txt")
	os.Create("backup.txt")
}

 func addNumber(b int){
 	primary,_ := os.OpenFile("primary.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0744)
 	 _, err := primary.WriteString(fmt.Sprintf("%d\n", b+1))
 	 fmt.Println(b+1)
 	 if err != nil{
 	 	log.Fatal(err)
 	 }
 	 time.Sleep(time.Second)
 }

 func addToBackupAndGetI() int{
 	_, err := os.Stat("primary.txt")
 	    if os.IsNotExist(err) {
 	       	 openPrimaryAndBackupFile(0)
 	   }
 	data, err := ioutil.ReadFile("primary.txt")
 	if err != nil {
 		log.Panicf("failed reading data from file: %s", err)
 	}
 	numbers := strings.Split(string(data), "\n")
 	num,_ := strconv.Atoi(numbers[len(numbers)-2])
 	backup,_ := os.OpenFile("backup.txt",os.O_WRONLY|os.O_CREATE, 0744)
 	backup.WriteString(fmt.Sprintf("%d\n", num))
 	if err != nil{
 		 log.Fatal(err)
 	}
 	return num
 }
 


func main() {
	backupnum := addToBackupAndGetI()

	for{
		if backupnum != addToBackupAndGetI(){
			break
		}
		backupnum += 1
		time.Sleep(time.Second)
		//maybe find a better solution
	}
	cmd := exec.Command("cmd", "/C", "start", "cmd", "go", "run", "main.go") //the second cmd can be changed with powershell
	err := cmd.Run()
	if err != nil { //getNewPrimaryandBackup()
		log.Fatal(err)
	}

	for {
		addToBackupAndGetI()
		openPrimaryAndBackupFile(addToBackupAndGetI())
		addNumber(addToBackupAndGetI())
	}

}