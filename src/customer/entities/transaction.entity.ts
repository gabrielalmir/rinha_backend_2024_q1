// type Transaction struct {
// 	ID          uint64    `gorm:"primaryKey;autoIncrement;column:id;index"`
// 	Amount      int64     `gorm:"column:valor;not null"`
// 	Type        string    `gorm:"column:tipo;not null"`
// 	Description string    `gorm:"column:descricao;not null"`
// 	CreatedAt   time.Time `gorm:"column:criado_em;not null"`
// 	CustomerID  uint64    `gorm:"column:cliente_id;not null;index"`
// }


interface TransactionProps {
    id?: number;
    amount: number;
    type: string;
    description: string;
    createdAt: Date;
    customerId: number;
}

export class Transaction {
    constructor(private props: TransactionProps) { }

    get id(): number {
        return this.props.id!;
    }

    get amount(): number {
        return this.props.amount;
    }

    get type(): string {
        return this.props.type;
    }

    get description(): string {
        return this.props.description;
    }

    get createdAt(): Date {
        return this.props.createdAt;
    }

    get customerId(): number {
        return this.props.customerId;
    }
}
