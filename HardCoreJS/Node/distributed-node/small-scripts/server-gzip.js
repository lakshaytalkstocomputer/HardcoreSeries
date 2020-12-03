const zlib = require("zlib");
const http = require("http");
const fs = require("fs");

http.createServer((request, response) => {
    // Read a file
    const raw = fs.createReadStream(__dirname + '/test-html.html');
    // find out which content encoding by reading header from request
    const acceptEncoding = request.headers['accept-encoding'] || "";
    // set content type header in response
    response.setHeader('Content-Type', 'text/plain');
    
    // compress the data depending on the content type if any
    if (acceptEncoding.includes('gzip')) {
        console.log("Encoding Type Found :", acceptEncoding);
        response.setHeader('COntent-Encoding', 'gzip');
        raw.pipe(zlib.createGzip()).pipe(response);        
    } else {
        console.log("No Encoding");
        raw.pipe(response);
    }   
}).listen(process.env.PORT || 3002);
