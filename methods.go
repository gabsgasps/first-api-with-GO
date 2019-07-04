package main

import "fmt"

type ff struct {
	peso, idade int
}

func (r *ff) calculo() int {
	return r.idade + r.peso
}

func (r ff) opa() int {
	return r.peso - r.idade
}

func main() {
	jj := ff{peso: 70, idade: 18}

	fmt.Println(jj.calculo())

	kk := &jj

	fmt.Println(kk.opa())

	kk.idade = 90

	fmt.Println(jj.idade)

}
