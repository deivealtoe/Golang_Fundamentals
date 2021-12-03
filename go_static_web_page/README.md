O módulo ```net/http``` é responsável por nos fornecer os métodos de mapeamento e gerenciamento de requisições e a função ```main``` será responsável a manter a aplicação de pé!

A função ```main```, como nos clássicos programas em ```C```, será o ponto de partida da execução do programa.

Dentro do ```main```, é utilizado as funções do pacote ```net/http```.

Uma função bem interessante do pacote para subir a aplicação é a ```ListenAndServe```, ela é responsável por especificar em qual porta rodará a aplicação e como lidar com o servidor da aplicação.

Para lidar com as requisições, é preciso utilizar um outro método do pacote ```net/http``` dentro da função ```main```, o ```HandleFunc```.

Como será responsável por lidar com as requisições e respostas do servidor, a função ```handler``` deve receber dois parâmetros, um do tipo ```http.ResponseWriter``` e outro do tipo ```*http.Request```.

Para servir a página estática, a função ```http.ServeFile``` pode ser usada.

O ```http.ServeFile``` é responsável por dispor um arquivo, ou seja, simplesmente servir uma página web utilizando as variáveis responsáveis por escrever a resposta ```http``` e também o request.
