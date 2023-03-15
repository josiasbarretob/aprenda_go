package main

import "fmt"
type veiculo struct{
	portas int
	cor string
}
type caminhonete struct{
	veiculo
	quatroRodas bool
}
type sedan struct{
	veiculo
	modeloLuxo bool
}

func main(){
	l200 := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor: "Preta",
		},
		quatroRodas: true,
	}
	cerato := sedan{
		veiculo: veiculo{
			portas: 4,
			cor: "vermelho",
		},
		modeloLuxo: true,
	}
	fmt.Println(l200)
	fmt.Println(cerato)
	fmt.Println("Cor da L200 é:",l200.cor)
	fmt.Println("Cor do Cerato é:",cerato.cor)
} 