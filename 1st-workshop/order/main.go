package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var (
	port int

	mysqlHost     string
	mysqlUser     string
	mysqlPassword string
	mysqlDatabase string

	emailUser     string
	emailPassword string
)

func main() {
	flag.IntVar(&port, "port", 9000, "http server port")
	flag.StringVar(&mysqlHost, "mysql_host", "localhost", "mysql host")
	flag.StringVar(&mysqlUser, "mysql_user", "root", "mysql user")
	flag.StringVar(&mysqlPassword, "mysql_password", "root-is-not-used", "mysql password")
	flag.StringVar(&mysqlDatabase, "mysql_database", "workshop_order", "mysql database")
	flag.StringVar(&emailUser, "email_user", "username@example.com", "email user")
	flag.StringVar(&emailPassword, "email_password", "secret", "email password")
	flag.Parse()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase)
	fmt.Printf("connecting to mysql on %s\n", dataSourceName)
	sqlDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	repo := NewMysqlRepo(sqlDB)
	emailService := NewEmailService(emailUser, emailPassword)
	router := mux.NewRouter()

	router.HandleFunc("/order", func(w http.ResponseWriter, req *http.Request) {
		if req.Body == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Empty request body"))
			return
		}
		var order Order
		err := json.NewDecoder(req.Body).Decode(&order)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		validatedOrder, err := order.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		lastInsertedID, err := repo.Store(validatedOrder)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		lastInsertedOrder, err := repo.ResolveByOrderID(lastInsertedID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		go func() {
			subject := fmt.Sprintf("Order: %s", lastInsertedOrder.OrderNumber)
			body := fmt.Sprintf(`
				Hello %s,

				<p>You just created an order in our website, please find the detail below.</p>

				<p>Order number: %s</p>
				<p>Total: %d</p>

				<p>Regards,</p>
				<p>Fundamental Go Workshop</p>
			`, lastInsertedOrder.FullName, lastInsertedOrder.OrderNumber, lastInsertedOrder.Total)
			err := emailService.Send(lastInsertedOrder.Email, subject, body)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Email for order %s just sent to %s\n", lastInsertedOrder.OrderNumber, "maderakateja@gmail.com")
			}
		}()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(lastInsertedOrder)
	}).Methods("POST")

	router.HandleFunc("/order/{orderId}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		orderIDString, ok := vars["orderId"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("orderId couldn't be found"))
			return
		}
		orderID, err := strconv.Atoi(orderIDString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		order, err := repo.ResolveByOrderID(int64(orderID))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	}).Methods("GET")

	router.HandleFunc("/order/{orderId}/html", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		orderIDString, ok := vars["orderId"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("orderId couldn't be found"))
			return
		}
		orderID, err := strconv.Atoi(orderIDString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		order, err := repo.ResolveByOrderID(int64(orderID))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		t := template.Must(template.ParseFiles("orderDetail.html"))
		t.Execute(w, order)
	})

	portString := fmt.Sprintf(":%d", port)
	log.Printf("listening on port %s", portString)
	http.ListenAndServe(portString, router)
}
