package main

import "fmt"

func main(){
    //cria slice s vazio com 3 posições
    s := make([]string, 3)
    fmt.Println("emp:", s)

    
    //adciona  às posições 0,1,2 respectivamente as letras a,b,c 
    s[0]="a"
    s[1]="b"
    s[2]="c"
    
    //printa o slice completo na tela e posteriormente apenas a posições 2
    fmt.Println("set:",s)
    fmt.Println("get:", s[2])
    
    //checa o tamanho do vetor
    fmt.Println("len:",len(s))
    
    //append adciona a ultima posição de s as letras passadas como argumento
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)
    
    //cria slice C com o mesmo tamanho de s
    c := make([]string, len(s))
    //copy copia de s para c os elementos
    copy(c,s)
    fmt.Println("cpy:",c)

    //copia os elementos nas posições de 2 a 5 para l e então imprime
    l:=s[2:5]
    fmt.Println("sli",l)

    

    l = s[:5]
    fmt.Println("sl2:",l)
    
    l= s[2:]
    fmt.Println("sl3:",l)

    t :=[]string{"g","h","i"}
    fmt.Println("dcl:",t)



    twoD := make([][]int,3)
    for i := 0; i < 3; i++{
        innerLen := i + i
        twoD[i] = make([]int,innerLen)
        for j := 0; j < innerLen; j++{
           twoD[i][j] = i + j
        }
    
    }

    fmt.Println("2d: ",twoD)
}
