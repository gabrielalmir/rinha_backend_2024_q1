import { describe, expect, it } from "vitest";
import { Transaction } from "./transaction.entity";

describe('create a transaction', () => {
    it('should create a transaction', () => {
        const transaction = new Transaction({
            amount: 0,
            createdAt: new Date(),
            customerId: 0,
            description: '',
            type: ''
        })

        expect(transaction).toBeInstanceOf(Transaction);
    });
})
