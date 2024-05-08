import { CustomerRepository } from "../repositories/customer.repository";

export class CustomerService {
    constructor(
        private readonly customerRepository: CustomerRepository,
    ) { }

    async getStatement(id: number) {
        return await this.customerRepository.getStatement(id);
    }
}
