import http from 'k6/http';
import { check } from 'k6';
import faker from "https://cdnjs.cloudflare.com/ajax/libs/Faker/3.1.0/faker.min.js";
//
// export let options = {
//     stages: [
//         { duration: '10s', target: 200},
//     ],
// };

export default function() {

    const url = 'http://nginx:4000/api/generate';

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
    check(response, {
        "is status 502": r => r.status === 502
    });
    check(response, {
        "is status 500": r => r.status === 500
    });
    check(response, {
        "is status 503": r => r.status === 503
    });

}