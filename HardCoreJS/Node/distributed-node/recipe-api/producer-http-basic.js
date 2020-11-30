const server = require("fastify")();
const HOST   = process.env.HOST || '127.0.0.1';
const PORT   = process.env.PORT || 4000;

console.log(`worker pid=${process.pid}`);

server.get('/recipe/:id', async(req, reply) => {
    console.log(`worker request pid=${process.pid}`);
    const id = Number(req.params.id);

    if(id !== 42){
        reply.statusCode = 404;
        return {
            error: 'not_found'
        }
    }
    return {
        producer_pid : process.pid,
        recipe : {
            id, name : "Chicken Tikka Masala",
            step     : "Throw it in pot!",
            ingrediants :[
                {
                    id       : 1,
                    name     : "Chicken",
                    quantity : "1 lb",
                },
                {
                    id       : 2,
                    name     : "Sauce",
                    quantity : "2 cups",
                }
            ]
        }
    };
});

server.listen(PORT, HOST, () => {
    console.log(`Producer Running at http://${HOST}:${PORT}`);
});
