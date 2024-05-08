import { TransactionDTO } from "../dtos/dtos";

export interface CustomerProps {
    id?: number;
    limit: number;
    balance: number;
    createdAt?: Date;
    transactions: TransactionDTO[];
}

export class Customer {
    constructor(private props: CustomerProps) { }

    get id(): number {
        return this.props.id!;
    }

    get limit(): number {
        return this.props.limit;
    }

    get balance(): number {
        return this.props.balance;
    }

    get createdAt(): Date {
        return this.props.createdAt!;
    }

    get transactions(): TransactionDTO[] {
        return this.props.transactions;
    }
}
