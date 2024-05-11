import { CustomerStatementDTO } from "../dtos/dtos";
import { CustomerService } from "../services/customer.service";

export class CustomerController {
    constructor(
        private readonly customerService: CustomerService,
    ) { }

    async getStatement(id: number): Promise<CustomerStatementDTO> {
        return await this.customerService.getStatement(id)
    }

    async createTransaction(id: number, valor: number, tipo: string, descricao: string) {
        return await this.customerService.createTransaction(id, valor, tipo, descricao)
    }
}
