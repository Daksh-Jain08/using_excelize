package main

import(
	"fmt"
	"strings"
	"github.com/xuri/excelize/v2"
	"encoding/json"
	"os"
	"io/ioutil"
)

type Meal struct {
	Day   string   `json:"day"`
	Date  string   `json:"date"`
	Meal  string   `json:"meal"`
	Items []string `json:"items"`
}

func (m *Meal) print_details() {
	fmt.Printf("Day: %v\n", m.Day)
	fmt.Printf("Date: %v\n", m.Date)
	fmt.Printf("Meal: %v\n", m.Meal)
	fmt.Printf("Items: %v\n", strings.Join(m.Items, ", "))
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

	colNum := -1
	for i := 0; i<7; i++{
		if rows[0][i] == day {
			colNum = i
			break
		}
	}
	if colNum == -1{
		fmt.Println("Please enter a valid date.")
		return nil, nil
	}

	mealRow := -1
	for i,row := range rows {
		cellValue := row[colNum]
		if cellValue == meal{
			mealRow = i
			break
		} 
	}
	if mealRow == -1{
		fmt.Println("Please enter a valid meal.")
		return nil, nil
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

	colNum := -1
	for i := 0; i<7; i++{
		if rows[0][i] == day {
			colNum = i
			break
		}
	}
	if colNum == -1{
		fmt.Println("Please enter a valid date.")
		return 0, nil
	}

	mealRow := -1
	for i,row := range rows {
		cellValue := row[colNum]
		if cellValue == meal{
			mealRow = i
			break
		} 
	}
	if mealRow == -1{
		fmt.Println("Please enter a valid date.")
		return 0, nil
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

	colNum := -1
	for i := 0; i<7; i++{
		if rows[0][i] == day {
			colNum = i
			break
		}
	}
	if colNum == -1{
		fmt.Println("Please enter a valid date.")
		return false, nil
	}

	mealRow := -1
	for i,row := range rows {
		cellValue := row[colNum]
		if cellValue == meal{
			mealRow = i
			break
		} 
	}
	if mealRow == -1{
		fmt.Println("Please enter a valid date.")
		return false, nil
	}

	for i := mealRow + 1;i < len(rows); i++ {
		value := rows[i][colNum]
		if value == item {
			return true, nil
		}
	}

	return false, nil
}

func convert_to_json(file string, choice string) ([]map[string]interface{}){
	f, err := excelize.OpenFile("Sample-Menu.xlsx")
    if err != nil {
        fmt.Println(err)
    }
	defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()

    cols, err := f.GetCols("Sheet1")
    if err != nil {
        fmt.Println(err)
    }

    var jsonData []map[string]interface{}

	meals := [3]string{"BREAKFAST", "LUNCH", "DINNER"}

	for _, col := range cols{
		
		mealNum := 0
		row := 3
		for i:=0;i<3;i++{
			var items []string
			for ;row < len(col); row++{
				if col[row] != "" {
					if col[row] == col[0] {
						row = row+2;
						break
					}

					items = append(items, col[row])
				}
			}
			jsonData = append(jsonData, map[string]interface{}{
					"day": col[0],
					"date": col[1],
					"meal": meals[mealNum],
					"items": items,
				},
			)
			mealNum++;
		}
	}

	if choice == "create_file" {
		jsonFile, err := os.Create("mess_menu.json")
		if err != nil {
			fmt.Println(err)
			return jsonData
		}
		defer func() {
			if err := jsonFile.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		encoder := json.NewEncoder(jsonFile)
		if err := encoder.Encode(jsonData); err != nil {
			fmt.Println(err)
			return jsonData
		}

		fmt.Println("The json file has been created with the name of mess_menu.json in the current working folder.")
	}
	return jsonData
}

func create_structs()([]Meal){
	var meals []Meal
	data, err := ioutil.ReadFile("mess_menu.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return meals
	}

	err = json.Unmarshal(data, &meals)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return meals
	}

	return meals
}

func main(){
	file := "Sample-Menu.xlsx"

	choice := -1
	for choice != 0{
		fmt.Scanln(&choice)

		var day,meal string
		if choice > 0 && choice < 4{
			fmt.Printf("Type the day of the week: ")
			fmt.Scanln(&day)
			day = strings.ToUpper(day)
			
			fmt.Printf("Type the meal of the day: ")
			fmt.Scanln(&meal)
			meal = strings.ToUpper(meal)
		}

		
		switch choice {
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
					println("")
				} else{
					fmt.Println(itemNum)
				}
			case 3:
				var item string
				fmt.Printf("Type the item you want to check: ")
				_, err := fmt.Scanf("%q", &item)
				if err != nil {
					fmt.Println("Error reading input:", err)
					break
				}

				item = strings.ToUpper(item)
				is, err := is_item_in_meal(file, day, meal, item)
				if err == nil && is {
					fmt.Println("Yes")
				} else {
					fmt.Println("No")
				}
				fmt.Scanln()
				break
			case 4:
				_ = convert_to_json(file, "create_file")
			case 5:
				meals := create_structs()

				for _,meal := range meals{
					meal.print_details()
				}
			default:
				fmt.Println("Please enter a valid choice.")
		}
	}
	fmt.Println("You have terminated the program!")

}