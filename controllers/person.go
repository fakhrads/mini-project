package controllers

import (
	"fmt"
	"formative-15/database"
	"formative-15/repository"
	"formative-15/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPerson(c *gin.Context) {
	var (
		result gin.H
	)

	fmt.Println(database.DbConntection)
	persons, err := repository.GetAllPerson(database.DbConntection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": persons,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertPerson(c *gin.Context) {
	var person structs.Person

	err := c.ShouldBindJSON(&person)

	if err != nil {
		panic(err)
	}

	err = repository.InsertPerson(database.DbConntection, person)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Person",
	})
}

func UpdatePerson(c *gin.Context) {
	var person structs.Person
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&person)
	if err != nil {
		panic(err)
	}

	person.ID = int64(id)

	err = repository.UpdatePerson(database.DbConntection, person)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Person",
	})
}

func DeletePerson(c *gin.Context) {
	var person structs.Person
	id, err := strconv.Atoi(c.Param("id"))

	person.ID = int64(id)

	err = repository.DeletePerson(database.DbConntection, person)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Person",
	})
}
