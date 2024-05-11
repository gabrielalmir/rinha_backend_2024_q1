export class CustomerTransactionBadRequest extends Error {
    constructor() {
        super('Invalid transaction type');
        this.name = 'CustomerTransactionBadRequest';
    }
}
