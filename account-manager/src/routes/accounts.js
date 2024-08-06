async function routes(fastify, options) {
    fastify.get('/accounts', async (request, reply) => {
        const accounts = await fastify.prisma.account.findMany();
        reply.send(accounts);
    });

    fastify.post('/accounts', async (request, reply) => {
        const { userId, type } = request.body;
        const account = await fastify.prisma.account.create({
            data: { userId, type }
        });
        reply.send(account);
    });
}

module.exports = routes;
