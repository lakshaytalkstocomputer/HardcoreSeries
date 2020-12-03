const fetch = require("node-fetch");

(async () => {
    const req = await fetch("http://localhost:3000",{
        method : "GET",
        headers : {
            'Content-Type' : 'applciation/json',
            'User-Agent' : `nodejs/${process.version}`,
            'Accept': 'applciation/json'
        }
        });
        
    const payload = await req.json();
    console.log("Payload :", payload);
})();
