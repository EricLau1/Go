'use strict';

const path = require("path");
const protoLoader = require('@grpc/proto-loader');
const grpc = require('grpc');

const PROTO_PATH = path.join('../pb', 'messages.proto');
const SERVER_ADDR = 'localhost:50000'; // Servidor onde está o função

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {});
const HelloService = grpc.loadPackageDefinition(packageDefinition).HelloService;

const client = new HelloService(SERVER_ADDR, grpc.credentials.createInsecure());

function main() {
    client.sayHello({Name: 'John Doe'}, (err, res) => {
        if (err) {
            console.log(err);
            return;
        }
        console.log(res.Message);
    });
}

main();