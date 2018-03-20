package apis

import (
	"strconv"
	"log"
	"fmt"
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
	."rest/models"
)

func AddPersonApi(c *gin.Context){
	firstName:=c.Request.FormValue("first_name")
	lastName:=c.Request.FormValue("last_name")

	person:=Person{FirstName:firstName,LastName:lastName}
	rows,err:=person.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("insert person Id {}",rows)
	msg:=fmt.Sprintf("insert successful %d",rows)
	c.JSON(http.StatusOK,gin.H{
		"msgCode":http.StatusOK,
		"msg":msg,
	})
}

func GetPersonsApi(c *gin.Context) {
	person:=Person{}
	persons,err:=person.GetPersons()
	if err!=nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK,gin.H{
		"persons":persons,
	})
}

func GetPersonApi(c *gin.Context) {
	cid:=c.Param("id")
	id,err:=strconv.Atoi(cid)
	person :=Person{Id:id}
	err = person.GetPerson()
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusOK,gin.H{
			"person":nil,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"person":person,
	})
}

func UpdatePersonApi(c *gin.Context) {
	cid:=c.Param("id")
	id,err:=strconv.Atoi(cid)
	person := Person{Id:id}
	err = c.Bind(&person)
	if err != nil{
		log.Fatalln(err)
	}

	ra,err:=person.UpdatePerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg:=fmt.Sprintf("Update person %d successful %d",person.Id,ra)
	c.JSON(http.StatusOK,gin.H{
		"msg":msg,
	})
}

func DeletePersonApi(c *gin.Context) {
	cid:=c.Param("id")
	id,err:=strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	person:=Person{Id:id}

	ra,err := person.DeletePerson()
	if err !=nil{
		log.Fatalln(err)
	}
	msg:=fmt.Sprintf("Delete person %d successful %d",id,ra)
	c.JSON(http.StatusOK,gin.H{
		"msg":msg,
	})
}