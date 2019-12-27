package route

import (
	"controle_pessoal_de_financas/API/v1/config"
	"controle_pessoal_de_financas/API/v1/controller"
	"net/http"
)

// Route é uma estrutura que representa cada rota do sistema. É composto por Name(nome), Method(método-POST, GET,...), Pattern(Padrão da rota, ex: /login/{usuario}), HandlerFunc(handler de pacote controller para a rota, com as regras e retornos específicos), Auth(Verdadeiro se na rota será obrigatório a passagem de TOKEN em anexo para acessá-la)
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Auth        bool
}

// Routes é um tipo que especifica o conjunto/lista de rotas no sistema(slice de Rotas)
type Routes []Route

// MyRoutes é uma variável que representa todas as rotas da API. Deve ser incluídas novas rotas caso necessário. Para informar o Method e o Pattern, deve ser alimentado primeiramente a variável Rotas em pacote config
var MyRoutes = Routes{
	Route{
		"API",
		config.Rotas["API"].Tipo,
		config.Rotas["API"].Rota,
		controller.API,
		false,
	},
	Route{
		"Index",
		config.Rotas["Index"].Tipo,
		config.Rotas["Index"].Rota,
		controller.Index,
		true,
	},
	Route{
		"Login",
		config.Rotas["Login"].Tipo,
		config.Rotas["Login"].Rota,
		controller.Login,
		false,
	},
	Route{
		"TokenValido",
		config.Rotas["TokenValido"].Tipo,
		config.Rotas["TokenValido"].Rota,
		controller.TokenValido,
		true,
	},
	Route{
		"PessoaIndex",
		config.Rotas["PessoaIndex"].Tipo,
		config.Rotas["PessoaIndex"].Rota,
		controller.PessoaIndex,
		true,
	},
	Route{
		"PessoaShow",
		config.Rotas["PessoaShow"].Tipo,
		config.Rotas["PessoaShow"].Rota,
		controller.PessoaShow,
		true,
	},
	Route{
		"PessoaShowAdmin",
		config.Rotas["PessoaShowAdmin"].Tipo,
		config.Rotas["PessoaShowAdmin"].Rota,
		controller.PessoaShowAdmin,
		true,
	},
	Route{
		"PessoaCreate",
		config.Rotas["PessoaCreate"].Tipo,
		config.Rotas["PessoaCreate"].Rota,
		controller.PessoaCreate,
		true,
	},
	Route{
		"PessoaRemove",
		config.Rotas["PessoaRemove"].Tipo,
		config.Rotas["PessoaRemove"].Rota,
		controller.PessoaRemove,
		true,
	},
	Route{
		"PessoaAlter",
		config.Rotas["PessoaAlter"].Tipo,
		config.Rotas["PessoaAlter"].Rota,
		controller.PessoaAlter,
		true,
	},
	Route{
		"PessoaEstado",
		config.Rotas["PessoaEstado"].Tipo,
		config.Rotas["PessoaEstado"].Rota,
		controller.PessoaEstado,
		true,
	},
	Route{
		"PessoaAdmin",
		config.Rotas["PessoaAdmin"].Tipo,
		config.Rotas["PessoaAdmin"].Rota,
		controller.PessoaAdmin,
		true,
	},
	Route{
		"TipoContaIndex",
		config.Rotas["TipoContaIndex"].Tipo,
		config.Rotas["TipoContaIndex"].Rota,
		controller.TipoContaIndex,
		true,
	},
	Route{
		"TipoContaShow",
		config.Rotas["TipoContaShow"].Tipo,
		config.Rotas["TipoContaShow"].Rota,
		controller.TipoContaShow,
		true,
	},
	Route{
		"TipoContaCreate",
		config.Rotas["TipoContaCreate"].Tipo,
		config.Rotas["TipoContaCreate"].Rota,
		controller.TipoContaCreate,
		true,
	},
	Route{
		"TipoContaRemove",
		config.Rotas["TipoContaRemove"].Tipo,
		config.Rotas["TipoContaRemove"].Rota,
		controller.TipoContaRemove,
		true,
	},
	Route{
		"TipoContaAlter",
		config.Rotas["TipoContaAlter"].Tipo,
		config.Rotas["TipoContaAlter"].Rota,
		controller.TipoContaAlter,
		true,
	},
	Route{
		"TipoContaEstado",
		config.Rotas["TipoContaEstado"].Tipo,
		config.Rotas["TipoContaEstado"].Rota,
		controller.TipoContaEstado,
		true,
	},
	Route{
		"ContaIndex",
		config.Rotas["ContaIndex"].Tipo,
		config.Rotas["ContaIndex"].Rota,
		controller.ContaIndex,
		true,
	},
	Route{
		"ContaShow",
		config.Rotas["ContaShow"].Tipo,
		config.Rotas["ContaShow"].Rota,
		controller.ContaShow,
		true,
	},
	Route{
		"ContaCreate",
		config.Rotas["ContaCreate"].Tipo,
		config.Rotas["ContaCreate"].Rota,
		controller.ContaCreate,
		true,
	},
	Route{
		"ContaRemove",
		config.Rotas["ContaRemove"].Tipo,
		config.Rotas["ContaRemove"].Rota,
		controller.ContaRemove,
		true,
	},
	Route{
		"ContaAlter",
		config.Rotas["ContaAlter"].Tipo,
		config.Rotas["ContaAlter"].Rota,
		controller.ContaAlter,
		true,
	},
	Route{
		"ContaEstado",
		config.Rotas["ContaEstado"].Tipo,
		config.Rotas["ContaEstado"].Rota,
		controller.ContaEstado,
		true,
	},
}
