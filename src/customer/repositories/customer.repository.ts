import { prisma } from "../../config/db"
import { TransactionDTO } from "../dtos/dtos"
import { CustomerTransactionBadRequest } from "../errors/customer-bad-request"
import { CustomerNotFound } from "../errors/customer-not-found"

export class CustomerRepository {
    async getStatement(id: number) {
        const customer = await prisma.customers.findFirst({
            where: { customer_id: id },
        })

        if (!customer) {
            throw new CustomerNotFound()
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
            throw new CustomerNotFound()
        }

        const signedAmount = transaction.type === 'c' ? transaction.amount : -transaction.amount

        const isValidTransaction = (
            transaction.type === 'd' || transaction.type === 'c')
            && customer.customer_balance + signedAmount <= customer.customer_limit

        if (!isValidTransaction) {
            throw new CustomerTransactionBadRequest()
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
