async function routes(fastify, options) {
    fastify.post('/send', async (request, reply) => {
        const transaction = request.body;
        processTransaction(transaction)
            .then(result => reply.send(result))
            .catch(error => reply.status(500).send({ error }));
    });

    fastify.post('/withdraw', async (request, reply) => {
        const transaction = request.body;
        processTransaction(transaction)
            .then(result => reply.send(result))
            .catch(error => reply.status(500).send({ error }));
    });
}

function processTransaction(transaction) {
    return new Promise((resolve, reject) => {
        console.log('Transaction processing started for:', transaction);

        setTimeout(() => {
            console.log('transaction processed for:', transaction);
            resolve(transaction);
        }, 30000);
    });
}

module.exports = routes;
