const server = require("fastify")();
const fetch  = require("node-fetch");

const HOST   = process.env.HOST   || '127.0.0.1';
const PORT   = process.env.PORT   || 3000;
const TARGET = process.env.TARGET || '127.0.0.1:4000';

server.get('/', async() => {
    const req = await fetch(`http://${TARGET}/recipes/42`);
    const producerData  = await req.json()

    return {
        consumerPid  : process.pid,
        producerData : producerData
    };
});

server.listen(PORT, HOST, () => {
    console.log(`Consumer running at pid: ${process.pid} http://${HOST}:${PORT}`);
});
