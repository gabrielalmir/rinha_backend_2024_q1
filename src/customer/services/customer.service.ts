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
            total += transaction.transaction_amount;
            transactionsDTO.push({
                description: transaction.transaction_description,
                type: transaction.transaction_type,
                amount: transaction.transaction_amount
            })
        }

        return { total, statementDate, transactions: transactionsDTO }
    }
}
