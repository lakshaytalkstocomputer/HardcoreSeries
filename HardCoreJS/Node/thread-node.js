const fs = require("fs");

// Node.JS Reads /etc/passwd, It is scheduled by libuv.
fs.readFile('/etc/passwd',
    (err, data) => {
        // Onece the file is done reading, libuv pases the result to the V8 event loop.
        if(err){
            throw err;
        }
        else{
            console.log(data);
        }
    }
);

// NodeJs runs a callback in a new stack, It's schduled by V8.
setImmediate(
    () => {
        // One's the previous stack is completed a new stack is created and it prints a message.
        console.log("This runs while file is being read!");
    }
);
