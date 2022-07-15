"use strict";

const sharp = require("sharp");

// sharp.concurrency(1);
// sharp.cache(false);

async function main() {
    console.log(`Running test ${process.argv[2]}`);

    switch (process.argv[2]) {
        case "resize":
            resize();
            break;
        default:
            throw new Error("Unknown test");
    }
    
}

async function resize() {
    let iterations = parseInt(process.argv[3]);

    for (let i = 0; i < iterations; i++) {
        console.log(`Step ${i}`)
        await resizeStep();
    }

}

async function resizeStep() {
    await sharp("images/image50.jpg").resize({width: 5000}).toFile("out/resize.jpg");
}

main();