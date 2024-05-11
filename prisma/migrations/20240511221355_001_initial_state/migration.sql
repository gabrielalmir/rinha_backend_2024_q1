-- CreateTable
CREATE TABLE "customers" (
    "customer_id" SERIAL NOT NULL,
    "customer_limit" INTEGER NOT NULL,
    "customer_balance" INTEGER NOT NULL,
    "created_at" TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "customers_pkey" PRIMARY KEY ("customer_id")
);

-- CreateTable
CREATE TABLE "transactions" (
    "transaction_id" SERIAL NOT NULL,
    "transaction_amount" INTEGER NOT NULL,
    "transaction_type" VARCHAR(255) NOT NULL,
    "transaction_description" VARCHAR(255) NOT NULL,
    "transaction_date" TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "transaction_customer" INTEGER NOT NULL,

    CONSTRAINT "transactions_pkey" PRIMARY KEY ("transaction_id")
);

-- CreateIndex
CREATE INDEX "customer_id_index" ON "customers"("customer_id");

-- CreateIndex
CREATE INDEX "transaction_customer_index" ON "transactions"("transaction_customer");

-- CreateIndex
CREATE INDEX "transaction_id_index" ON "transactions"("transaction_id");

-- AddForeignKey
ALTER TABLE "transactions" ADD CONSTRAINT "transactions_transaction_customer_fkey" FOREIGN KEY ("transaction_customer") REFERENCES "customers"("customer_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
