
export interface CustomerDTO {
    limit: number
    balance: number
}

export interface CustomerStatementDTO {
    total: number;
    statementDate: Date;
    transactions: TransactionDTO[];
}

export interface TransactionDTO {
    description: string;
    type: string;
    amount: number
}
