package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Guest struct {
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
}

func MZRverify() gin.HandlerFunc {
    return func(c *gin.Context) {
        var reqBody Guest 
        var guests []Guest 
        var response bool = false

        // Will be Hotel API services use in this part of the project.
        // But currently we only use fakeData.
        guests = append(guests, Guest{
            FirstName: "DAVIDS",
            LastName: "SNYDERS",
        })
        guests = append(guests, Guest{
            FirstName: "DAVID",
            LastName: "SNYDER",
        })

        if err := c.BindJSON(&reqBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error request Data": err.Error()})
            return
        }

		for i := 0; i < len(guests); i++ {
			if reqBody.FirstName ==  guests[i].FirstName && 
            reqBody.LastName == guests[i].LastName {
                response = true
                c.JSON(http.StatusOK, gin.H{"response": response})

                return
			}
        }

        response = false
        c.JSON(http.StatusOK, gin.H{"response": response})
    }
}