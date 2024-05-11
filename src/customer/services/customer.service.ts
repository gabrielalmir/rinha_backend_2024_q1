import { CustomerRepository } from "../repositories/customer.repository";

export class CustomerService {
    constructor(
        private readonly customerRepository: CustomerRepository,
    ) { }

    async getStatement(id: number) {
        const { transactions, statementDate } = await this.customerRepository.getStatement(id);

        let total = 0;
        const transactionsDTO = []

        for (const transaction of transactions) {
            const signedAmount = transaction.transaction_type === 'c' ? transaction.transaction_amount : -transaction.transaction_amount;

            total += signedAmount;

            transactionsDTO.push({
                description: transaction.transaction_description,
                type: transaction.transaction_type,
                amount: transaction.transaction_amount
            })
        }

        return { total, statementDate, transactions: transactionsDTO }
    }

    async createTransaction(id: number, valor: number, tipo: string, descricao: string) {
        return await this.customerRepository.createTransaction(id, { amount: valor, type: tipo, description: descricao })
    }
}
