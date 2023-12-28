'use client'

import { LevelRequest } from '/gen/game_pb.js';
import { ColorSortApiClient } from '/gen/game_grpc_web_pb.js';

export default function Page() {
    grpcCall();
    return <h1>Hello, Next.js!</h1>
}

function grpcCall() {
    var service = new ColorSortApiClient('http://localhost:50051');
    const request = new LevelRequest();
    request.setLevel(1);

    service.getGameLevel(request, {}, console.log);

}