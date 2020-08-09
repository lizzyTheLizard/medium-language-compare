const { Client } = require('pg');

const clientConfig = {
    user: "postgres",
    database: "postgres",
    password: "postgres",
    port: 5432,
    host: "postgres",
};

exports.connect = async function() {
    console.log('Try to connect to DB', clientConfig)
    for (var index = 0; index < 10; index++) {
        const client = new Client(clientConfig);
        try {
            await client.connect()
            console.log('Connected to database')
            return client;
        } catch (err) {
            console.log('Failed to connect to DB', err)
            lastError = err;
        }
        await new Promise(resolve => setTimeout(resolve, 1000));
    }
    throw lastError;
};
