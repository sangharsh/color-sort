"use client"

import { NewLevelPlayRequest, PourRequest, ResetRequest, UndoRequest } from '/gen/game_pb.js';
import { ColorSortApiClient } from '/gen/game_grpc_web_pb.js';

const API_URL = 'http://localhost:8080';
const METADATA = { 'colorsort-userid': 'abc123' };
const SERVICE = new ColorSortApiClient(API_URL);

export function NewLevel(levelId, callback) {
    const request = new NewLevelPlayRequest();
    request.setId(levelId);

    const processResponse = (err, response) => {
        if (err) {
            console.log("err:", err, "response: ", response);
            return;
        }
        callback(response);
    };

    SERVICE.newLevel(request, METADATA, processResponse);
}

export function Pour(src, dst, callback) {
    const req = new PourRequest();
    req.setSrc(src);
    req.setDst(dst);

    const processResponse = (err, response) => {
        if (err) {
            console.log("err:", err, "response: ", response);
            return;
        }
        callback(response.getLevel());
    };
    SERVICE.pour(req, METADATA, processResponse);
}

export function Reset(callback) {
    const req = new ResetRequest();
    SERVICE.reset(req, METADATA, callback);
}
export function Undo(callback) {
    const req = new UndoRequest();
    SERVICE.undo(req, METADATA, callback);

}