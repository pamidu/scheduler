package main

import (
	"fmt"
	"time"
	"github.com/streadway/amqp"
	"log"
	"encoding/json"
	"strconv"
	"os"

)
type ScheduleObject struct{
	Timestamp string
	OperationData map[string]interface{}
	ControlData map[string]interface{}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func Request(){
	fmt.Println("Request")
	objectsset:=getFakeObjects()
	//fmt.Println(objectsset)


conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()


fmt.Println("test 01 ")
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for ob:=range objectsset{
		
	t := time.Now().Local()
	//fmt.Println(t)

	abc:=t.Format("20060102150405")

 	currentTimeInt, err := strconv.Atoi(abc)
    if err != nil {
          fmt.Println(err)
        
    }

s:=ScheduleObject{Timestamp:"20141120094352"}
fmt.Println(objectsset[ob].Timestamp,"test",currentTimeInt,s,"\n\n")
Inttimestamp,_:=strconv.Atoi(objectsset[ob].Timestamp)
timedifference:=Inttimestamp-currentTimeInt
fmt.Println(timedifference)

//fmt.Println(abc.Sub(objectsset[ob].Timestamp))

if timedifference>=0 && timedifference<=30{
		dataset, _ := json.Marshal(objectsset[ob])

	
		
		//fmt.Println("dfgh",s.Timestamp,"21212112")
		body :=dataset
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
		})
}else if timedifference<=0{
	//fmt.Println("0000000000")
	//break
	os.Exit(99)

}
}
		
	//}
	failOnError(err, "Failed to publish a message")










}


func main (){

/*form := "20141120154115"
    //form := "2014-11-20 15:44:45"
    t2, e := time.Parse(form,"2015-11-20 15:44:45")
fmt.Println(t2,"timexxx",e)

*/
t := time.Now().Local()
//fmt.Println(t.Format("20141120171013"),"asdfghjkl;")
if t!=time.Now(){

}
	c := time.Tick(1* time.Second)
	for now := range c {
		if now!=time.Now() {
			
		}
	    //fmt.Printf("Time: %v \n", now)
	    go Request()
	}


}



func getFakeObjects()(objects []ScheduleObject){
	objects = make([]ScheduleObject,1)
	//fmt.Println("getFakeObjects method run ")

	for index,_ := range objects{
		//t := time.Now().Local()
		//fmt.Println(t.Format("20060102150405"))
		

		var tmpOperationData map[string]interface{}
		tmpOperationData = make(map[string]interface{})
		tmpOperationData["a"] = 1
		tmpOperationData["b"] = 2

		var tmpControlData map[string]interface{}
		tmpControlData = make(map[string]interface{})
		tmpControlData["a"] = 3
		tmpControlData["b"] = 4
		tmpControlData["c"] = 5
		tmpControlData["d"] = 6
		tmpControlData["e"] = 7
		tmpControlData["f"] = 8

		t := "20141121132919"//time.Now().Local()
		objects[index].Timestamp=t//.Format("20060102150405")
		objects[index].OperationData = tmpOperationData
		objects[index].ControlData = tmpControlData

	}



	return objects

}

