import { prisma } from "../../config/db"
import { TransactionDTO } from "../dtos/dtos"

export class CustomerRepository {
    async getStatement(id: number) {
        const customer = await prisma.customers.findFirst({
            where: { customer_id: id },
        })

        if (!customer) {
            throw new Error('Customer not found')
        }

        const transactions = await prisma.transactions.findMany({
            where: { transaction_customer: id },
        })

        const total = transactions.reduce((acc, transaction) => {
            return acc + transaction.transaction_amount
        }, 0)

        return {
            total,
            statementDate: new Date(),
            transactions,
        }
    }

    async createTransaction(customerId: number, transaction: TransactionDTO) {
        const customer = await prisma.customers.findFirst({
            where: { customer_id: customerId },
        })

        if (!customer) {
            throw new Error('Customer not found')
        }

        if ((transaction.type !== 'd' && transaction.type !== 'c') && customer.customer_balance + transaction.amount <= customer.customer_limit) {
            throw new Error('Invalid transaction type')
        }

        const newTransaction = await prisma.transactions.create({
            data: {
                transaction_description: transaction.description,
                transaction_type: transaction.type,
                transaction_amount: transaction.amount,
                transaction_customer: customerId,
            },
        })

        return newTransaction
    }
}
