import fastify from "fastify";
import zod from 'zod';
import { env } from "./config/env";

import { prisma } from "./config/db";
import { CustomerController } from "./customer/controllers/customer.controller";
import { CustomerTransactionBadRequest } from "./customer/errors/customer-bad-request";
import { CustomerNotFound } from "./customer/errors/customer-not-found";
import { CustomerRepository } from "./customer/repositories/customer.repository";
import { CustomerService } from "./customer/services/customer.service";

const app = fastify();

const customerController = new CustomerController(
    new CustomerService(new CustomerRepository())
);

app.get("/clientes/:id/extrato", async (request, reply) => {
    const schema = zod.object({
        id: zod.string()
    });

    const { id } = schema.parse(request.params);

    const { success, data, error } = await customerController.getStatement(+id)

    if (!success) {
        if (error instanceof CustomerNotFound) {
            return reply.status(404).send();
        }
        return reply.status(500).send();
    }

    return {
        total: data?.total,
        statementDate: data?.statementDate,
        transactions: data?.transactions,
    };
})

app.post("/clientes/:id/transacoes", async (request, reply) => {
    const requestSchema = zod.object({ id: zod.string() });
    const bodySchema = zod.object({
        valor: zod.number(),
        tipo: zod.string(),
        descricao: zod.string(),
    });

    const { id } = requestSchema.parse(request.params);
    const { valor, tipo, descricao } = bodySchema.parse(request.body);

    const { success, error } = await customerController.createTransaction(+id, valor, tipo, descricao)

    if (!success) {
        if (error instanceof CustomerNotFound) {
            return reply.status(404).send();
        }
        else if (error instanceof CustomerTransactionBadRequest) {
            return reply.status(400).send();
        }
        return reply.status(500).send();
    }

    return reply.status(201).send();
})

app.listen({ port: env.PORT }, async (err, address) => {
    if (err) {
        console.error(err);
        await prisma.$disconnect();
        process.exit(1);
    }
    console.log(`Server listening at ${address}`);
    await prisma.$connect();
});

export { app };
