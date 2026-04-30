# 📌 Convenções e Padrões da Linguagem Go

> Go é conhecida por sua filosofia de "menos é mais", priorizando a legibilidade e a simplicidade sobre a complexidade sintática.  
> Seguir as convenções da comunidade (muitas vezes referenciadas como *idiomatic Go*) é fundamental para manter um código sustentável.

---

## 1. Naming e Visibilidade (Variáveis e Funções)
Em Go, a visibilidade de um identificador é definida pela letra inicial. Não existem palavras-chave como `public` ou `private`.

- **PascalCase**: Identificadores iniciados com letra maiúscula são **exportados** (públicos para outros pacotes).  
- **camelCase**: Identificadores iniciados com letra minúscula são **não-exportados** (privados ao pacote).  
- **Acrônimos**: Devem ser mantidos em caixa alta ou baixa de forma consistente.  
  - Correto: `serveHTTP`, `HTTPRequest`  
  - Incorreto: `serveHttp`, `HttpRequest`  
- **Nomes Curtos**: Go favorece nomes curtos para escopos reduzidos.  
  - Em loops: `i`, `j`  
  - Para receptores de struct: `f` para `File`, nunca `this` ou `self`.

---

## 2. Structs
As `structs` são a base para a organização de dados.

- **Composição sobre Herança**: Go não possui herança. Usa-se *embedding* para promover campos e métodos de uma struct para outra.  
- **Construtores**: Por convenção, structs complexas são inicializadas por funções que começam com `New`.  
  - Exemplo: `user.New()`  
- **Tags**: Use tags de struct para metadados (como `json:"user_id"`), mantendo consistência com APIs externas sem quebrar o padrão interno camelCase.

---

## 3. Interfaces
Interfaces em Go são implementadas implicitamente (*structural typing*).

- **Naming**: Interfaces de um único método geralmente terminam em `-er`.  
  - Exemplo: `Reader`, `Writer`, `Closer`.  
- **Pequenas e Poderosas**: Prefira interfaces pequenas. A biblioteca padrão é famosa por interfaces de 1–2 métodos.  
- **Aceite interfaces, retorne structs**: Funções devem aceitar interfaces para flexibilidade, mas retornar tipos concretos para clareza.

---

## 4. Tratamento de Erros
Go trata erros como valores de retorno, não com `try/catch`.

- **Verificação imediata**: A convenção é verificar o erro logo após a chamada.  
  ```go
  if err != nil {
      return err
  }
  ```
- **Guard Clauses**: Prefira retornos precoces para evitar aninhamentos excessivos.  
- **Erros contextuais**: Use `fmt.Errorf` ou `errors.Join` para adicionar contexto.  
  ```go
  return fmt.Errorf("falha ao salvar usuário: %w", err)
  ```

---

## ✅ Resumo
- **Naming**: Maiúscula = exportado, minúscula = interno.  
- **Structs**: Composição > herança, construtores com `New`.  
- **Interfaces**: Pequenas, terminando em `-er`, aceitam interfaces e retornam structs.  
- **Erros**: Tratados como valores, verificação imediata e guard clauses.