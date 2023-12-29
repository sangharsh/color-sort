'use client'

import { LevelRequest } from '/gen/game_pb.js';
import { ColorSortApiClient } from '/gen/game_grpc_web_pb.js';

export default function Page() {
    grpcCall();
    return <h1>Hello, Next.js!</h1>
}

function grpcCall() {
    var service = new ColorSortApiClient('http://localhost:8080');
    const request = new LevelRequest();
    request.setLevel(1);

    const processResponse = (err, response) => {
        console.log("err:", err, "response: ", response);
        console.log("level:", response.getLevel());
        console.log("tubes:", response.getTubesList());
        const tubes = response.getTubesList();
        const t0 = tubes[0];
        console.log("t0:", t0);
        const colors = t0.getColorsList()
        console.log("colors:", colors);
        const c0 = colors[0];
        console.log("c0:", c0);
    };

    service.getGameLevel(request, {}, processResponse);

}