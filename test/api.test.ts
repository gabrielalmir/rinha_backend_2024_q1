
import { beforeAll, describe, it } from "vitest";

describe("Valida status da API", () => {
    beforeAll(async () => {
        const { app } = await import("../src/index");
        await app.listen({ port: 3000 });
    });

    it("Liste extrato de um cliente", async () => {
        const response = await fetch("http://localhost:3000/clientes/1/extrato");
        return response.status === 200;
    });

    it("Liste extrato de um cliente não existente", async () => {
        const response = await fetch("http://localhost:3000/clientes/6/extrato");
        return response.status === 404;
    });

    it("Criar uma transação a um cliente existente", async () => {
        const response = await fetch("http://localhost:3000/clientes/1/transacoes", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                valor: 220,
                tipo: "c",
                descricao: "descricao",
            }),
        });
        return response.status === 201;
    });

    it("Criar uma transação a um cliente não existente", async () => {
        const response = await fetch("http://localhost:3000/clientes/6/transacoes", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                valor: 1000,
                tipo: "c",
                descricao: "descricao",
            }),
        });
        return response.status === 404;
    });

    it("Criar uma transação inválida a um cliente", async () => {
        const response = await fetch("http://localhost:3000/clientes/1/transacoes", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                valor: 100000,
                tipo: "c",
                descricao: "descricao",
            }),
        });
        return response.status === 400;
    })
})
