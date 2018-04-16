package main

import (
	"fmt"
	"math"
)
func main(){
	var (
		a, a2 [5][9]float64 //Начальная матрица и Пересчитанная по методу Гаусса
		bc [5] float64 //Массив коэффициентов b/c
		i, j, im, jm int
		chek bool
		err int
		el float64
		c string
		)

	fmt.Println("Данная программа предназначена для решения задачи симплекс методом \nВсего 4 вида товаров!\nПродолжить? (y/n)")
	fmt.Scanln(&c)
	if c[0]=='y' {
	
		name := map[int]string{
			0: "сталь",
			1: "чугун",
			2: "бронза",
			3: "медь",
		}
		fmt.Println("ok")
		for i=0; i<5; i++ {
			for j:=0; j<4;j++{
				if i==4 {
					fmt.Println("введите цену",name[j]) // коэффициенты с
				}else{
					fmt.Println("для тов", i+1, "введите кол.",name[j]) // коэфициенты a[i][j]
				}
				fmt.Scanln(&a[i][j])

				//заполняем матрицу единицами по диагонали
				if i==j {
					a[i][j+4] = 1
				}else{
					a[i][j+4] = 0
				} 
			}
			if i<4 {
				fmt.Println("(максимально возможный)") //коэффициенты b
				fmt.Scanln(&a[i][8])
			}
		} 
   //    | a11 a12 a13 a14 1   0   0   0   b1 |
   //    | a21 a22 a23 a24 0   1   0   0   b2 |
   // A  | a31 a32 a33 a34 0   0   1   0   b3 |
   //    | a41 a42 a43 a44 0   0   0   1   b4 |
   //    | c1  c2  c3  c4  0   0   0   0   0  |  

		//выводим получившуюся матрицу
		for i=0; i<5; i++ {
			fmt.Printf("%7.2f",a[i][0:4])
			fmt.Print(" ")
			fmt.Printf("%7.2f",a[i][4:8])
			fmt.Printf("%7.2f",a[i][8])
			fmt.Println()
		}

		err = 0
		chek = false
		for err < 11 && !chek {
			err++
			// находим разрешающий столбец (макс среди C)
			el = 0
			jm = 1
			for j=0; j<9; j++ {
				if a[4][j]*a[4][j] > el*el {
					el = a[4][j]
					jm = j
				}
			}
			// находим b[i]/c[i,jm]
      		for i=0; i<5; i++ {
      			if a[i][jm] == 0 {
      				bc[i]=1.7e36 //типа очень бльшое число
      			} else {
      				bc[i]=a[i][8]/a[i][jm]
      			}
      		}
		    //находим разрешающюю строку (мин. среди b[i]/c[i,разрешающий столбец])
		    el = bc[0]
		    im = 1
		    for i=0; i<5; i++ {
		      if el > bc[i] {
		         im = i
		         el = bc[i]
		      	}
		    }
		    //a[im,jm]- разрешающий элемент
			//Производим пресчет по методу Гаусса
			if a[im][jm] == 0 {
				fmt.Println("разрешающий элемент = 0")
			} else {
				for i=0; i<5; i++ { 
				 	for j=0; j<9; j++ {
				 		if i == im {
				 			a2[i][j] = a[i][j]/a[im][jm]
				 		} else {
				 			a2[i][j] = a[i][j]-(a[i][jm]*a[im][j])/a[im][jm]
				 		}
					}
				}
			}	
		 	//перебрасываем после подсчетов данные из массива А2 в А
			a = a2
			fmt.Println("im =",im,"jm =",jm)
			//выводим получившуюся матрицу
			for i=0; i<5; i++ {
				fmt.Printf("%7.2f",a[i][0:4])
				fmt.Print(" ")
				fmt.Printf("%7.2f",a[i][4:8])
				fmt.Printf("%7.2f",a[i][8])
				fmt.Println()
			}
			fmt.Println()

			chek = true
			for j=0; j<9; j++ {
				if a[4][j]>0 {
					chek = false
				}
			}	
		}
		if err==10{
			fmt.Println("repeat until false ;)")
		} else {
			fmt.Println("Z max =",-a[4][8]);
			for j=0; j<5; j++ {
				for i=0; i<5; i++ {
					a[4][j]=a[4][j]+math.Abs(a[i][j])
				}
			} 
			for j=0; j<5; j++ {
				if a[4][j]==1 {
					for i=0; i<5; i++{
						if a[i][j]==1 {
							fmt.Println("X",j,"=",a[i][8])
						}
					}
				}
			}
		}
	}


}