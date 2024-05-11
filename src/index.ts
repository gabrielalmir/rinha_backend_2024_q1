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

    return await customerController.getStatement(+id)
        .catch((err) => {
            console.log(err);
            if (err instanceof CustomerNotFound) {
                return reply.status(404).send();
            }
            return reply.status(500).send();
        })
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

    return await customerController.createTransaction(+id, valor, tipo, descricao)
        .then(() => reply.status(201).send())
        .catch((err) => {
            if (err instanceof CustomerNotFound) {
                return reply.status(404).send();
            }
            else if (err instanceof CustomerTransactionBadRequest) {
                return reply.status(400).send();
            }
            return reply.status(500).send();
        })
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
