const fastify = require('fastify')({ logger: true });
const prisma = require('./prisma/client');
const accountRoutes = require('./routes/accounts');
const userRoutes = require('./routes/users');

fastify.register(require('@fastify/jwt'), {
    secret: 'supersecret'
});

fastify.register(userRoutes);
fastify.register(accountRoutes);

const start = async () => {
    try {
        await fastify.listen(3000, '0.0.0.0');
        fastify.log.info(`server listening on ${fastify.server.address().port}`);
    } catch (err) {
        fastify.log.error(err);
        process.exit(1);
    }
};
start();
