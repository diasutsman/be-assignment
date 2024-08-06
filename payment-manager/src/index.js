const fastify = require('fastify')({ logger: true });
const prisma = require('./prisma/client');
const transactionRoutes = require('./routes/transactions');

fastify.register(require('@fastify/jwt'), {
    secret: 'supersecret'
});

fastify.register(transactionRoutes);

const start = async () => {
    try {
        await fastify.listen(3001, '0.0.0.0');
        fastify.log.info(`server listening on ${fastify.server.address().port}`);
    } catch (err) {
        fastify.log.error(err);
        process.exit(1);
    }
};
start();
