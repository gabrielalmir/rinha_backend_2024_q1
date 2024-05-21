import 'dotenv/config'
import zod from 'zod'

// Define environment schema
const envSchema = zod.object({
    NODE_ENV: zod.string().default('production'),
    PORT: zod.number().default(3000),
    DATABASE_URL: zod.string().default('file:./dev.db'),
})

export type Env = zod.infer<typeof envSchema>
export const env = envSchema.parse(process.env) as Env
