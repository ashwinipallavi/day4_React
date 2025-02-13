package main

import (
	//"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Flight struct {
	Id          string  `json:"id"`
	Number      string  `json:"number"`
	ArilineName string  `json:"ariline_name"`
	Source      string  `json:"source"`
	Destination string  `json:"destination"`
	Capacity    int     `json:"capacity"`
	Price       float32 `json:"price"`
}

func readAllFlights(c *gin.Context) {
	flights := []Flight{
		{Id: "1001", Number: "AI 845",
			ArilineName: "Air India", Source: "Mumbai",
			Destination: "Abu Dhabi", Capacity: 180,
			Price: 15000.0},
		{Id: "1002", Number: "AI 846",
			ArilineName: "Air India", Source: "Abu Dhabi",
			Destination: "Mumbai", Capacity: 180,
			Price: 15000.0},
	}
	c.JSON(http.StatusOK, flights)
}

func readFlightById(c *gin.Context) {
	id := c.Param("id")
	flight := Flight{Id: id, Number: "AI 845",
		ArilineName: "Air India", Source: "Mumbai",
		Destination: "Abu Dhabi", Capacity: 180,
		Price: 15000.0}
	c.JSON(http.StatusOK, flight)
}

func createFlight(c *gin.Context) {
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Error." + err.Error()})
		return
	}
	createdFlight := Flight{Id: "1001", Number: "AI 845",
		ArilineName: "Air India", Source: "Mumbai",
		Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0}
	c.JSON(http.StatusCreated,
		gin.H{"message": "Flight Created Successfully.",
			"flight": createdFlight})
}

func updateFlight(c *gin.Context) {
	id := c.Param("id") //
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Error." + err.Error()})
		return
	}
	updatedFlight := Flight{Id: id, Number: "AI 845", //
		ArilineName: "Air India", Source: "Mumbai",
		Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0}
	c.JSON(http.StatusOK,
		gin.H{"message": "Flight Updated Successfully.", //
			"flight": updatedFlight}) //
}

func deleteFlight(c *gin.Context) {
	//id := c.Param("id")
	c.JSON(http.StatusOK,
		gin.H{"message": "Flight Deleted Successfully."})
}
func main() {
	//router
	r := gin.Default()
	//cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//routes | API Endpoints
	r.GET("/flights", readAllFlights)
	r.GET("/flights/:id", readFlightById)
	r.POST("/flights", createFlight)
	r.PUT("/flights/:id", updateFlight)
	r.DELETE("/flights/:id", deleteFlight)
	//server
	r.Run(":8080")
}
