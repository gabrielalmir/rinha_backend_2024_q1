-- CreateTable
CREATE TABLE "customers" (
    "customer_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "customer_limit" INTEGER NOT NULL,
    "customer_balance" INTEGER NOT NULL,
    "created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- CreateTable
CREATE TABLE "transactions" (
    "transaction_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "transaction_amount" INTEGER NOT NULL,
    "transaction_type" TEXT NOT NULL,
    "transaction_description" TEXT NOT NULL,
    "transaction_date" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "transaction_customer" INTEGER NOT NULL,
    CONSTRAINT "transactions_transaction_customer_fkey" FOREIGN KEY ("transaction_customer") REFERENCES "customers" ("customer_id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

-- CreateIndex
CREATE INDEX "customer_id_index" ON "customers"("customer_id");

-- CreateIndex
CREATE INDEX "transaction_customer_index" ON "transactions"("transaction_customer");

-- CreateIndex
CREATE INDEX "transaction_id_index" ON "transactions"("transaction_id");
