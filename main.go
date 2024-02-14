package main

import(
	"fmt"
	"strings"
	"github.com/xuri/excelize/v2"
)

type Meal struct {
	day   string
	date  string
	meal  string
	items []string
}

func (m *Meal) PrintDetails() {
	fmt.Println("DAY: " + m.day)
	fmt.Println("Date: " + m.date)
	fmt.Println("Meal: " + m.meal)
	fmt.Printf("Items: %v", m.items)
	fmt.Println()
}

func get_items(file,day,meal string) ([]string,error) {
	
	f,err := excelize.OpenFile(file)
	if err != nil { 
		return nil,err
	}

	defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()

	rows,err := f.GetRows("Sheet1")
	if err != nil {
		return nil,err
	}

	colNum := 0
	for i := 0; i<7; i++{
		if rows[0][i] == day {
			colNum = i
			break
		}
	}

	mealRow := 0
	for i,row := range rows {
		cellValue := row[colNum]
		if cellValue == meal{
			mealRow = i
			break
		} 
	}

	var items []string
	for i := mealRow + 1;i < len(rows); i++ {
		item := rows[i][colNum]
		if item == day {
			break
		}

		items = append(items, item)
	}

	return items, nil
}

func get_number_of_items(file,day,meal string) (int,error) {
	
	f,err := excelize.OpenFile(file)
	if err != nil { 
		return 0,err
	}

	defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()

	rows,err := f.GetRows("Sheet1")
	if err != nil {
		return 0,err
	}

	colNum := 0
	for i := 0; i<7; i++{
		if rows[0][i] == day {
			colNum = i
			break
		}
	}

	mealRow := 0
	for i,row := range rows {
		cellValue := row[colNum]
		if cellValue == meal{
			mealRow = i
			break
		} 
	}

	num := 0
	for i := mealRow + 1;i < len(rows); i++ {
		item := rows[i][colNum]
		if item == day || item == "" {
			break
		}
		num++
	}

	return num, nil
}

func is_item_in_meal(file,day,meal,item string) (bool,error) {
	
	f,err := excelize.OpenFile(file)
	if err != nil { 
		return false,err
	}

	defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()

	rows,err := f.GetRows("Sheet1")
	if err != nil {
		return false,err
	}

	colNum := 0
	for i := 0; i<7; i++{
		if rows[0][i] == day {
			colNum = i
			break
		}
	}

	mealRow := 0
	for i,row := range rows {
		cellValue := row[colNum]
		if cellValue == meal{
			mealRow = i
			break
		} 
	}

	for i := mealRow + 1;i < len(rows); i++ {
		value := rows[i][colNum]
		if value == item {
			return true, nil
		}
	}

	return false, nil
}

func main(){
	file := "Sample-Menu.xlsx"

	choice := 0
	for {
		fmt.Scanln(&choice)
		fmt.Printf("Type the day of the week: ")
		var day string
		fmt.Scanln(&day)
		day = strings.ToUpper(day)
		
		fmt.Printf("Type the meal of the day: ")
		var meal string
		fmt.Scanln(&meal)
		meal = strings.ToUpper(meal)
		
		switch choice {
			case 0:
				fmt.Println("You have terminated the program!")
				return
			case 1:
				items, err := get_items(file, day, meal)
				if err != nil {
					fmt.Println(err)
				}
				for i := range items {
					fmt.Println(items[i])
				}
			case 2:
				itemNum, err := get_number_of_items(file, day, meal)
				if err != nil {
					fmt.Println(err)
				}
				if itemNum == 0{
					println("some error occurred!")
				} else{
					fmt.Println(itemNum)
				}
			case 3:
				item := ""
				fmt.Printf("Typoe the item you want to check: ")
				fmt.Scanln(&item)
				is, err := is_item_in_meal(file, day, meal, item)
				if err == nil && is {
					fmt.Printf("Yes")
				} else {
					fmt.Printf("No")
				}
			default:
				fmt.Println("Please enter a valid choice.")
		}
	}
}