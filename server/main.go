package main

import (
	"context" //**
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"           //**
	"go.mongodb.org/mongo-driver/bson/primitive" //**
	"go.mongodb.org/mongo-driver/mongo"          //**
	"go.mongodb.org/mongo-driver/mongo/options"  //**
)

// Config
var mongoUri string = "mongodb://localhost:27017"
var mongoDbName string = "ars_app_db"

// Database variables
var mongoClient *mongo.Client
var employeeCollection *mongo.Collection

// Employee struct (updated from Flight to Employee)
type Employee struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Position   string             `json:"position" bson:"position"`
	Department string             `json:"department" bson:"department"`
	Salary     float64            `json:"salary" bson:"salary"`
	Email      string             `json:"email" bson:"email"`
	Phone      string             `json:"phone" bson:"phone"`
}

// MongoDB connection function
func connectToMongo() {
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		fmt.Println("Mongo DB Connection Error!" + err.Error())
		return
	}
	name := "employees" // Update the collection name to "employees"
	employeeCollection = mongoClient.Database(mongoDbName).Collection(name)
}

// API Endpoints

// Create a new employee
func createEmployee(c *gin.Context) {
	var employee Employee
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	//
	employee.Id = primitive.NewObjectID()
	_, err := employeeCollection.InsertOne(context.TODO(), employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	//
	c.JSON(http.StatusCreated,
		gin.H{"message": "Employee Created Successfully", "employee": employee})
}

// Get all employees
func readAllEmployees(c *gin.Context) {
	//
	var employees []Employee
	cursor, err := employeeCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	defer cursor.Close(context.TODO())
	//
	employees = []Employee{}
	err = cursor.All(context.TODO(), &employees)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	//
	c.JSON(http.StatusOK, employees)
}

// Get employee by ID
func readEmployeeById(c *gin.Context) {
	id := c.Param("id")
	//
	employeeId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID.\n" + err.Error()})
		return
	}
	//
	var employee Employee
	err = employeeCollection.FindOne(context.TODO(), bson.M{"_id": employeeId}).Decode(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Employee Not Found."})
		return
	}
	//
	c.JSON(http.StatusOK, employee)
}

// Update an employee by ID
func updateEmployee(c *gin.Context) {
	id := c.Param("id")
	employeeId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID.\n" + err.Error()})
		return
	}
	//
	var oldEmployee Employee
	err = employeeCollection.FindOne(context.TODO(), bson.M{"_id": employeeId}).Decode(&oldEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Employee Not Found."})
		return
	}
	//
	var jbodyEmployee Employee
	err = c.BindJSON(&jbodyEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	oldEmployee.Salary = jbodyEmployee.Salary
	//
	_, err = employeeCollection.UpdateOne(context.TODO(),
		bson.M{"_id": employeeId},
		bson.M{"$set": bson.M{"salary": jbodyEmployee.Salary}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	//response
	c.JSON(http.StatusOK, gin.H{"message": "Employee Updated Successfully", "employee": oldEmployee})
}

// Delete an employee by ID
func deleteEmployee(c *gin.Context) {
	id := c.Param("id")
	employeeId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID.\n" + err.Error()})
		return
	}
	//
	var oldEmployee Employee
	err = employeeCollection.FindOne(context.TODO(), bson.M{"_id": employeeId}).Decode(&oldEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Employee Not Found."})
		return
	}
	//delete
	_, err = employeeCollection.DeleteOne(context.TODO(), bson.M{"_id": employeeId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	//response
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully."})
}

// Main function
func main() {
	//
	connectToMongo() // Connect to MongoDB
	// router
	r := gin.Default()
	// CORS configuration (Allow frontend URL)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// routes (API Endpoints)
	r.POST("/employees", createEmployee)       // Create employee
	r.GET("/employees", readAllEmployees)      // Read all employees
	r.GET("/employees/:id", readEmployeeById)  // Read employee by ID
	r.PUT("/employees/:id", updateEmployee)    // Update employee
	r.DELETE("/employees/:id", deleteEmployee) // Delete employee
	// server (run on port 8080)
	r.Run(":8081")
}
