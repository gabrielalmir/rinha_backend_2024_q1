
import { beforeAll, describe, expect, it } from "vitest";
import { z } from "zod";
import { prisma } from "../src/config/db";

describe("Valida status da API", () => {
    beforeAll(async () => {
        await prisma.transactions.deleteMany();
        await prisma.customers.deleteMany();
        await prisma.customers.createMany({
            data: [
                { customer_id: 1, customer_limit: 100000, customer_balance: 0 },
                { customer_id: 2, customer_limit: 80000, customer_balance: 0 },
                { customer_id: 3, customer_limit: 1000000, customer_balance: 0 },
                { customer_id: 4, customer_limit: 10000000, customer_balance: 0 },
                { customer_id: 5, customer_limit: 500000, customer_balance: 0 },
            ],
        });

        const { app } = await import("../src/index");
        await app.listen({ port: 3000 });
    });

    it("Liste extrato de um cliente", async () => {
        const response = await fetch("http://localhost:3000/clientes/1/extrato");
        return response.status === 200;
    });

    it("Extrato de um cliente com dados corretos", async () => {
        const response = await fetch("http://localhost:3000/clientes/1/extrato");
        const data = await response.json();
        const schema = z.object({
            total: z.number(),
            statementDate: z.string(),
            transactions: z.array(z.object({
                description: z.string(),
                type: z.string(),
                amount: z.number(),
            })),
        });
        const { statementDate, total, transactions } = schema.parse(data);

        expect(statementDate !== null && total !== null && transactions.length === 0).toBe(true);
    });

    it("Liste extrato de um cliente não existente", async () => {
        const response = await fetch("http://localhost:3000/clientes/6/extrato");
        expect(response.status === 404).toBe(true);
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

        expect(response.status === 201).toBe(true);
    });

    it("Verificar se o saldo do cliente foi atualizado", async () => {
        const response2 = await fetch("http://localhost:3000/clientes/1/extrato");
        const data = await response2.json();

        const schema = z.object({
            total: z.number(),
            statementDate: z.string(),
            transactions: z.array(z.object({
                description: z.string(),
                type: z.string(),
                amount: z.number(),
            })),
        });

        const { total } = schema.parse(data);

        expect(total === 220).toBe(true);
    })

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

        expect(response.status === 404).toBe(true);
    });

    it("Criar uma transação inválida a um cliente", async () => {
        const response = await fetch("http://localhost:3000/clientes/1/transacoes", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                valor: 100_000_000_000, // Valor maior que o limite
                tipo: "c",
                descricao: "descricao",
            }),
        });

        expect(response.status === 400).toBe(true);
    })
})
