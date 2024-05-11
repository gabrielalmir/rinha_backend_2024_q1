import { CustomerRepository } from "../repositories/customer.repository";

export class CustomerService {
    constructor(
        private readonly customerRepository: CustomerRepository,
    ) { }

    async getStatement(id: number) {
        return await this.customerRepository.getStatement(id)
    }

    async createTransaction(id: number, valor: number, tipo: string, descricao: string) {
        return await this.customerRepository.createTransaction(id, { amount: valor, type: tipo, description: descricao })
    }
}
