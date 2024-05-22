
# Rinha Backend 2024/Q1

Este é o repositório do projeto **Rinha Backend 2024/Q1**, uma aplicação backend desenvolvida com TypeScript e diversas ferramentas modernas para proporcionar uma estrutura robusta e escalável.

## Sumário

- [Instalação](#instalação)
- [Configuração](#configuração)
- [Scripts Disponíveis](#scripts-disponíveis)
- [Estrutura de Pastas](#estrutura-de-pastas)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Testes](#testes)
- [Contribuição](#contribuição)
- [Licença](#licença)

## Instalação

Para instalar as dependências do projeto, execute:

```bash
npm install
```

## Configuração

### Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto e adicione as seguintes variáveis de ambiente:

```env
DATABASE_URL=your_database_url
```

### Configuração do Docker

Para executar o projeto utilizando Docker, certifique-se de ter o Docker instalado em sua máquina e execute:

```bash
docker-compose up
```

## Scripts Disponíveis

No diretório do projeto, você pode executar:

### `npm run dev`

Roda a aplicação em modo de desenvolvimento.Abra [http://localhost:3000](http://localhost:3000) para ver no navegador.

### `npm run build`

Compila a aplicação para produção na pasta `dist`.

### `npm start`

Inicia a aplicação em produção a partir da pasta `dist`.

### `npm test`

Executa os testes utilizando o Vitest.

## Estrutura de Pastas

```plaintext
src/
├── @types/               # Definições de tipos
├── controllers/          # Controladores da aplicação
├── dtos/                 # Objetos de Transferência de Dados
├── errors/               # Classes de erros personalizados
├── repositories/         # Repositórios de dados
├── services/             # Serviços de negócio
├── .dockerignore         # Arquivo Docker ignore
├── .editorconfig         # Configurações do editor
├── .gitattributes        # Atributos Git
├── .gitignore            # Arquivo Git ignore
├── Dockerfile            # Dockerfile para criação da imagem Docker
├── compose.yaml          # Arquivo de configuração do Docker Compose
├── eslint.config.js      # Configuração do ESLint
├── package-lock.json     # Lockfile do npm
├── package.json          # Arquivo de configuração do npm
├── tsconfig.json         # Configuração do TypeScript
└── vite.config.ts        # Configuração do Vite
```

## Tecnologias Utilizadas

- **Node.js**: Ambiente de execução para JavaScript.
- **TypeScript**: Superconjunto de JavaScript que adiciona tipos estáticos.
- **Express**: Framework web para Node.js.
- **Prisma**: ORM para Node.js e TypeScript.
- **Docker**: Plataforma para desenvolvimento de aplicações em containers.
- **Vitest**: Framework de testes unitários.
- **ESLint**: Ferramenta para análise estática de código.
- **Vite**: Ferramenta de build e desenvolvimento.

## Testes

Para executar os testes, utilize o comando:

```bash
npm test
```

## Contribuição

Contribuições são bem-vindas! Por favor, abra uma issue ou envie um pull request.

1. Faça um fork do projeto.
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`).
3. Faça commit das suas alterações (`git commit -m 'Adiciona nova feature'`).
4. Faça push para a branch (`git push origin feature/nova-feature`).
5. Abra um pull request.
