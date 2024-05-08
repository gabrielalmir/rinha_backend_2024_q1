import fastify from "fastify";
import zod from 'zod';
import { env } from "./config/env";

import { CustomerController } from "./customer/controllers/customer.controller";
import { CustomerRepository } from "./customer/repositories/customer.repository";
import { CustomerService } from "./customer/services/customer.service";

const app = fastify();

const customerController = new CustomerController(
    new CustomerService(new CustomerRepository())
);

app.get("/clientes/:id/extrato", async (request) => {
    const schema = zod.object({
        id: zod.string()
    });

    const { id } = schema.parse(request.params);

    return await customerController.getStatement(+id);
})

app.listen({ port: env.PORT }, (err, address) => {
    if (err) {
        console.error(err);
        process.exit(1);
    }
    console.log(`Server listening at ${address}`);
});
