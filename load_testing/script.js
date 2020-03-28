import http from 'k6/http';
import { check } from 'k6';
import faker from "https://cdnjs.cloudflare.com/ajax/libs/Faker/3.1.0/faker.min.js";

export let options = {
    stages: [
        { duration: '10s'},
        { duration: '10s'},
    ],
};

export default function() {

    const url = 'http://elf-server:5000/api/generate';

    const payload = {
        originalUrl: faker.random.uuid(),
        expiresAfter: "45"
    };

    const params = {
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
    };

    let response = http.post(url, payload, params);

    check(response, {
        "is status 200": r => r.status === 200
    });

}