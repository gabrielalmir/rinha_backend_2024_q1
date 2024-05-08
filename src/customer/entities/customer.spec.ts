import { describe, expect, it } from 'vitest';
import { Customer } from './customer.entity';

describe('create a customer', () => {
    it('should create a customer', () => {
        const customer = new Customer({
            balance: 0,
            limit: 0,
        })

        expect(customer).toBeInstanceOf(Customer);
    });

    it('should create a customer with balance', () => {
        const customer = new Customer({
            balance: 100,
            limit: 0,
        })

        expect(customer.balance).toBe(100);
    });

    it('should create a customer with limit', () => {
        const customer = new Customer({
            balance: 0,
            limit: 100,
        })

        expect(customer.limit).toBe(100);
    });
})
