import {Client, ClientConfig} from 'pg'


const clientConfig: ClientConfig = {
    user: "postgres",
    database: "postgres",
    password: "postgres",
    port: 5432,
    host: "postgres",
};

export async function connect() : Promise<Client> {
    var lastError: Error = new Error();
    console.log('Try to connect to DB', clientConfig)
    for (var index = 0; index < 10; index++) {
        const client = new Client(clientConfig);
        try {
            await client.connect()
            console.log('Connected to database')
            return client;
        } catch (err) {
            lastError = err;
        }
        await new Promise(resolve => setTimeout(resolve, 1000));
    }
    throw lastError;
}
