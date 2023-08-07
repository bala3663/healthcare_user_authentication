package controllers

import (
	"Final-Project-gin_helthcare/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" //high performance HTTP(hypertext transfer protocol) web framework
	_ "github.com/go-sql-driver/mysql"
)

//get method

func Get_patients1() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/Patient_information")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Database connected")
		defer db.Close()
		results, err := db.Query("SELECT * FROM patient_info_2")
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()

		var output interface{}
		for results.Next() {
			var Patient_Id int
			var Name string
			var Address string
			var PhoneNo int
			var Disease string
			var Modifier string
			var Payer string
			var Bill_no int
			var Bill_Amount int
			var Paid_by_Payer int
			var Remaining_Amount int
			err = results.Scan(&Patient_Id, &Name, &Address, &PhoneNo, &Disease, &Modifier, &Payer, &Bill_no, &Bill_Amount, &Paid_by_Payer, &Remaining_Amount)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d %s %s %d %s %s %s %d %d %d %d ", Patient_Id, Name, Address, PhoneNo, Disease, Modifier, Payer, Bill_no, Bill_Amount, Paid_by_Payer, Remaining_Amount)
			c.IndentedJSON(200, "patients")
			c.JSON(http.StatusOK, gin.H{"": output})
		}

	}

}

// post method

func Post_patient1() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/Patient_information")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		var new_patient models.Info
		err = c.BindJSON(&new_patient)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf(`INSERT INTO patient_info_2 (Patient_Id, Name, Address, PhoneNo, Disease, Modifier, Payer, Bill_no, Bill_Amount, Paid_by_Payer) VALUES(%d,"%s","%s",%d,"%s","%s","%s",%d,%d,%d)`, new_patient.Patient_Id, new_patient.Name, new_patient.Address, new_patient.PhoneNo, new_patient.Disease, new_patient.Modifier, new_patient.Payer, new_patient.Bill_no, new_patient.Bill_Amount, new_patient.Paid_by_Payer)

		insert, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
		c.IndentedJSON(201, "Yes, values added!")
	}

}

// Put Method

func Update_patientinfo1() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/Patient_information")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("database connected")
		defer db.Close()
		var edit_patient_info models.Info
		err = c.BindJSON(&edit_patient_info)
		if err != nil {
			return
		}
		query := fmt.Sprintf("UPDATE patient_info_2 SET  Name = '%s', Address = '%s', PhoneNo ='%d', Disease = '%s',Modifier = '%s',Payer = '%s',Bill_no = '%d',Bill_Amount = '%d',Paid_by_Payer = '%d' WHERE Patient_Id = '%d' ", edit_patient_info.Name, edit_patient_info.Address, edit_patient_info.PhoneNo, edit_patient_info.Disease, edit_patient_info.Modifier, edit_patient_info.Payer, edit_patient_info.Bill_no, edit_patient_info.Bill_Amount, edit_patient_info.Paid_by_Payer, edit_patient_info.Patient_Id)
		_, err = db.Exec(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(201, "Yes, Successfully Updated")
	}
}

func Delete_Patient1() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/Patient_information")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("database connected")
		defer db.Close()
		var delete_patient_info models.Info
		err = c.BindJSON(&delete_patient_info)
		if err != nil {
			return
		}
		query := fmt.Sprintf(" DELETE FROM patient_info_2 WHERE Patient_Id=%d ", delete_patient_info.Patient_Id)
		_, err = db.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, "Successfully Deleted")
	}
}

func Search_patient_by_bill_id1() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/Patient_information")
		if err != nil {
			panic(err.Error())
		}
		var search_byid models.Info
		err = c.BindJSON(&search_byid)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf("SELECT * FROM patient_info_2 WHERE Bill_no='%d'", search_byid.Bill_no)
		results, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var Patient_Id int
			var Name string
			var Address string
			var PhoneNo int
			var Disease string
			var Modifier string
			var Payer string
			var Bill_no int
			var Bill_Amount int
			var Paid_by_Payer int
			var Remaining_Amount int
			err = results.Scan(&Patient_Id, &Name, &Address, &PhoneNo, &Disease, &Modifier, &Payer, &Bill_no, &Bill_Amount, &Paid_by_Payer, &Remaining_Amount)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d %s %s %d %s %s %s %d %d %d %d ", Patient_Id, Name, Address, PhoneNo, Disease, Modifier, Payer, Bill_no, Bill_Amount, Paid_by_Payer, Remaining_Amount)
			c.IndentedJSON(200, "patients")
			c.JSON(http.StatusOK, gin.H{"": output})
		}

	}

}
