package main

import (
	"fmt"
)

func main(){
	//Slices são os arrays dinamicos do go

	a := [3]int{1,2,3}
	fmt.Println(a)
	//um slice é declarado utilizando dois indices, minimo e maximo, separados por ':'
	fmt.Println(a[0:1])
	fmt.Println(a[0:2])
	//slices não armazenam dados, apenas descrevem uma seção de um array subjacente.
	//os indicies definem quais elementos do array serão selecionados.
	fmt.Println(a[0:3])
	fmt.Println(a[1:2])
	fmt.Println(a[1:3])
	//O indice máximo é exclusivo
	// fmt.Println(a[1:4]) indexOutOfBounds

	//modificar um elemento de um slice modifica o elemento do array subjacente

	nomes :=[3]string{"sapo","gato","ra"}
	//Cursiosidade: é necessário adicionar uma vírgula
	//após o último elemento de um array de strings caso
	//ele seja declarado com a seguinte notação:
	// teste :=[2]string{
	// 	"t1",
	// 	"t2"
	// }
	//descomente o codigo acima e veja o editor apontar um erro na linha onde esta t2

	fmt.Println("Array completo:", nomes)

	b := nomes[1:3]
	fmt.Println("slice b: ",b)

	b[0] = "XXX"
	fmt.Println("Array completo com um dos elementos alterados através do slice b: ",nomes)

	//slices diferentes construídos sobre o mesmo array subjacente verão as mudanças causadas por slices

	c := [3]int{1,2,3}

	c1 := c[0:3]
	c2 := c[0:3]

	fmt.Println("Array completo", c)
	fmt.Println("Slices c1 e c2", c1,c2)

	c2[0] = 5

	fmt.Println("Slice c2 com o primeiro elemento alterado: ",c2)
	fmt.Println("Array completo após alteração através do slice c2: ",c)
	fmt.Println("Slice c1 refletindo a alteração do slice c2: ", c1)

	//Ao utilizar slices, pode-se omitir os valores de indices mínimos e máximos e usar seus valores padrão
	//o valor padrão do indice minimo é zero e o valor padrão do indice máximo é o comprimento do array

	//o comprimento e a capacidade de um slice podem ser obtidos utilizando as funcoes len() e cap()

	d :=[3]int{22,33,44}
	d1 := d[:]
	fmt.Println("Capacidade e comprimento do slice d1: ",cap(d1), len(d1))

	//é possivel alterar o comprimento de um slice  refazendo o "slicing" sobre ele mesmo
	//isso é, se houve capacidade disponível:

	d1 = d1[:1]
	fmt.Println("Reduzindo o comprimento do slice d1, que agora tem comprimento de :", len(d1))

	d1 = d1[:len(d)]
	fmt.Println("Extendendo o comprimento do slice d1, que agora tem o comprimento de :", len((d1)))

	d1 = d1[1:]
	fmt.Println("Reduzindo a capacidade do slice d1, que agora tem capacidade de ",cap(d1)," e comprimento de", len(d1))

	//A linha abaixo causa um erro
	// d1 = d1[5:]
	// fmt.Println("Aumentando o comprimento do slice d1 além de sua capacidade: ", len(d1))

	//é possível criar um slice com a funcao make():

	e := make([]int,1,5)
	fmt.Printf("Retorno do metodo make: %d ; Um array comum: %d",e,a)
	fmt.Println(a,e)


	//para adicionar elementos a um slice basta usar o metodo append()
	f:= []int{1,2,3}
	fmt.Printf("Slice f: %d com endereço de memoria %p\n", f, &f)

	f = append(f, 4)

	fmt.Printf("Slice f após um append: %d com endereço de memoria %p \n", f, &f)

	//pode-se appendar mais de um valor por vez a um slice

	f = append(f, 6,7,8)
	fmt.Printf("Slice f com append multiplo: %d com endereço de memoria %p\n", f, &f)

	//pode-se também appendar um slice dentro do outro

	g :=[]int{9,10,11}
	f = append(f,g...) // necessario usar o "spread operator" para transformar o slice em uma lista de argumentos
	fmt.Printf("Slice f com um append de outro slice: %d com endereço de memoria %p\n", f, &f)
	












}