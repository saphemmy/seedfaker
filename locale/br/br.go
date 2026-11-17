// Package br provides the Brazil locale for seedfaker.
package br

import "github.com/saphemmy/seedfaker"

// Locale is Brazil: +55 phones, 8-digit CEP codes (#####-###).
var Locale = seedfaker.Locale{
	Code:            "pt_BR",
	Country:         "Brazil",
	FirstNames:      []string{"Miguel", "Alice", "Arthur", "Sophia", "Heitor", "Helena", "Davi", "Valentina", "Gabriel", "Laura", "Bernardo", "Isabella", "Lucas", "Manuela", "Matheus"},
	LastNames:       []string{"Silva", "Santos", "Oliveira", "Souza", "Rodrigues", "Ferreira", "Alves", "Pereira", "Lima", "Gomes", "Costa", "Ribeiro", "Martins", "Carvalho"},
	Titles:          []string{"Sr", "Sra", "Srta", "Dr", "Dra"},
	PhoneFormats:    []string{"+55 ## 9#### ####"},
	Cities:          []string{"Sao Paulo", "Rio de Janeiro", "Brasilia", "Salvador", "Fortaleza", "Belo Horizonte", "Manaus", "Curitiba", "Recife", "Porto Alegre"},
	States:          []string{"Sao Paulo", "Rio de Janeiro", "Minas Gerais", "Bahia", "Parana", "Rio Grande do Sul", "Pernambuco", "Ceara"},
	Streets:         []string{"Rua das Flores", "Avenida Paulista", "Rua XV de Novembro", "Avenida Brasil", "Rua Sete de Setembro", "Avenida Atlantica"},
	PostcodeFormats: []string{"#####-###"},
}

func init() { seedfaker.Register(Locale) }
