package main

import (
	"strconv"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
     
	output := ""

	input:=widget.NewLabel("")
    
	var Harr []string;
    
	Hstr:=""

	history:=widget.NewLabel(Hstr);
    
	isH:=false

	historyBtn:=widget.NewButton("History",func(){

          if isH{
             Hstr=""
		  }else{
			for i:=len(Harr)-1;i>=0;i--{
				Hstr+=Harr[i]
				Hstr+="\n"
			  }
		  }

		  isH = !isH

		  history.SetText(Hstr)
	})

	backBtn:=widget.NewButton("Back",func(){

		if len(output)>0{
			output = output[:len(output)-1]
		    input.SetText(output);
		}
	})

	clearBtn:=widget.NewButton("Clear",func(){
		output = ""
		input.SetText(output)
	})

	openBtn:=widget.NewButton("(",func(){
		output += "("
		input.SetText(output)
	})
    
	closeBtn:=widget.NewButton(")",func(){
		output += ")"
		input.SetText(output)
	})

	divideBtn:=widget.NewButton("/",func(){
		output += "/"
		input.SetText(output)
	})

	sevenBtn:=widget.NewButton("7",func(){
		output += "7"
		input.SetText(output)
	})

	eightBtn:=widget.NewButton("8",func(){
		output += "8"
		input.SetText(output)
	})

	nineBtn:=widget.NewButton("9",func(){
		output += "9"
		input.SetText(output)
	})


	mulBtn:=widget.NewButton("*",func(){
		output += "*"
		input.SetText(output)
	})

	fourBtn:=widget.NewButton("4",func(){
		output += "4"
		input.SetText(output)
	})

	fiveBtn:=widget.NewButton("5",func(){
		output += "5"
		input.SetText(output)
	})

	sixBtn:=widget.NewButton("6",func(){
		output += "6"
		input.SetText(output)
	})

	minusBtn:=widget.NewButton("-",func(){
		output += "-"
		input.SetText(output)
	})

	oneBtn:=widget.NewButton("1",func(){
		output += "1"
		input.SetText(output)
	})

	twoBtn:=widget.NewButton("2",func(){
		output += "2"
		input.SetText(output)
	})

	threeBtn:=widget.NewButton("3",func(){
		output += "3"
		input.SetText(output)
	})

	plusBtn:=widget.NewButton("+",func(){
		output += "+"
		input.SetText(output)
	})

	zeroBtn:=widget.NewButton("0",func(){
		output += "0"
		input.SetText(output)
	})

	dotBtn:=widget.NewButton(".",func(){
		output += "."
		input.SetText(output)
	})

	equalBtn:=widget.NewButton("=",func(){
		expression, err := govaluate.NewEvaluableExpression(output);

	    if err==nil{
			result, err := expression.Evaluate(nil);
			if err==nil{
				ans := strconv.FormatFloat(result.(float64),'f',-1,64)
                strToAppend := output+"="+ans
				Harr = append(Harr,strToAppend)
				output=ans
			}else{
				output = "error"
			}
		}else{
			output = "error"
		}

		input.SetText(output)
	})

	

	w.SetContent(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1, //1st row consist
			container.NewGridWithColumns(2, //two buttons
                 historyBtn,
				 backBtn,
			),

			

		     container.NewGridWithColumns(4, //next row consist four buttons
			    clearBtn,
			    openBtn,
			    closeBtn,
			    divideBtn,
	        ),

			container.NewGridWithColumns(4, //next row consist four buttons
			    nineBtn,
			    eightBtn,
			    sevenBtn,
			    mulBtn,
	        ),

			container.NewGridWithColumns(4, //next row consist four buttons
			    fourBtn,
				fiveBtn,
				sixBtn,
				minusBtn,
	        ),

			container.NewGridWithColumns(4, //next row consist four buttons
			    oneBtn,
				twoBtn,
				threeBtn,
				plusBtn,
	        ),

			container.NewGridWithColumns(2, //next row consist four buttons
			    container.NewGridWithColumns(2, //next row consist four buttons
			       zeroBtn,
				   dotBtn,
				),

				   equalBtn,
	        ),

		
		),
	),)

	w.ShowAndRun()
}
